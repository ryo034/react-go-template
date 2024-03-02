package middleware

import (
	"context"
	"database/sql"
	"fmt"

	fbDr "github.com/ryo034/react-go-template/apps/system/api/driver/firebase"

	"github.com/ryo034/react-go-template/apps/system/api/domain/me/provider"

	"github.com/go-faster/errors"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

	"github.com/google/uuid"

	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/core"

	authDr "github.com/ryo034/react-go-template/apps/system/api/driver/auth"

	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	firebase *fb.Firebase
	co       shared.ContextOperator
	dbp      core.Provider
	authDr   authDr.Driver
	fbDr     fbDr.Driver
}

func NewSecMiddleware(firebase *fb.Firebase, co shared.ContextOperator, dbp core.Provider, authDr authDr.Driver, fbDr fbDr.Driver) *Middleware {
	return &Middleware{firebase, co, dbp, authDr, fbDr}
}

// HandleBearer run only when the authentication method is Bearer
func (m *Middleware) HandleBearer(ctx context.Context, operationName string, t openapi.Bearer) (context.Context, error) {
	return m.checkAndSetToken(ctx, t.GetToken())
	//newCtx, err := m.checkAndSetToken(ctx, t.GetToken())
	//if err != nil {
	//	return nil, err
	//}
	//return m.authorize(newCtx, operationName)
}

// checkAndSetToken return JWT information in the key token in the argument context
func (m *Middleware) authorize(ctx context.Context, operationName string) (context.Context, error) {
	return ctx, nil
}

func (m *Middleware) checkAndSetToken(ctx context.Context, t string) (context.Context, error) {
	token, err := m.firebase.Auth.VerifyIDToken(ctx, t)
	if err != nil || token == nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
	}
	apUID, err := provider.NewUID(token.UID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
	}
	ctx = m.co.SetAuthProviderUID(ctx, apUID)
	var aID uuid.UUID
	err = uuid.Validate(token.UID)
	if err == nil {
		aID = uuid.MustParse(token.UID)
	} else {
		tkInClaim := token.Claims[fbDr.CustomClaimAccountIDKey]
		if tkInClaim != nil && tkInClaim != "" {
			if err = uuid.Validate(tkInClaim.(string)); err != nil {
				return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
			}
			aID = uuid.MustParse(tkInClaim.(string))
		} else {
			aID, err = m.authDr.FindAccountIDByAuthProviderUID(ctx, m.dbp.GetExecutor(ctx, true), apUID)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
				}
				// only OAuth provider
				providerInfo, ok := token.Claims["firebase"].(map[string]interface{})
				if !ok {
					return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
				}
				signInProvider, ok := providerInfo["sign_in_provider"].(string)
				if !ok {
					return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
				}
				switch signInProvider {
				case "google.com":
					tmpAID, _ := account.GenerateID()
					aID = tmpAID.Value()
				}
			}
			if err = m.fbDr.SetAccountIDToCustomClaim(ctx, account.NewIDFromUUID(aID)); err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Server Error"))
			}
		}
	}
	// Return JWT information in the key token in the argument context
	return m.co.SetClaim(m.co.SetUID(ctx, aID.String()), token.Claims), nil
}

//func adaptFirebaseUserInfoToAuthProviderUserInfo(ur *fbV4uth.UserInfo) (*provider.UserInfo, error) {
//	id, err := account.NewID(ur.UID)
//	if err != nil {
//		return nil, err
//	}
//	return &provider.UserInfo{
//		DisplayName: ur.DisplayName,
//		Email:       ur.Email,
//		PhoneNumber: ur.PhoneNumber,
//		PhotoURL:    ur.PhotoURL,
//		ProviderID:  ur.ProviderID,
//		UID:         id,
//	}, nil
//}

//func adaptFirebaseUserInfoListToAuthProviderUserInfoList(ur []*fbV4uth.UserInfo) ([]*provider.UserInfo, error) {
//	result := make([]*provider.UserInfo, 0, len(ur))
//	for _, u := range ur {
//		au, err := adaptFirebaseUserInfoToAuthProviderUserInfo(u)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, au)
//	}
//	return result, nil
//}

//func adaptFirebaseUserMetadataToAuthProviderUserMetadata(ur *fbV4uth.UserMetadata) *provider.UserMetadata {
//	return &provider.UserMetadata{
//		CreationTimestamp:    ur.CreationTimestamp,
//		LastLogInTimestamp:   ur.LastLogInTimestamp,
//		LastRefreshTimestamp: ur.LastRefreshTimestamp,
//	}
//}

//func adaptFirebaseUserEnrolledFactorsToAuthProviderUserEnrolledFactors(mfs []*fbV4uth.MultiFactorInfo) []*provider.MultiFactorInfo {
//	result := make([]*provider.MultiFactorInfo, 0, len(mfs))
//	for _, mf := range mfs {
//		id, err := account.NewID(mf.UID)
//		if err != nil {
//			return nil
//		}
//		result = append(result, &provider.MultiFactorInfo{
//			UID:                 id,
//			DisplayName:         mf.DisplayName,
//			EnrollmentTimestamp: mf.EnrollmentTimestamp,
//			FactorID:            mf.FactorID,
//			Phone:               &provider.PhoneMultiFactorInfo{PhoneNumber: mf.Phone.PhoneNumber},
//		})
//	}
//	return result
//}

//func adaptFirebaseUserToAuthProviderUser(ur *fbV4uth.UserRecord) (*provider.User, error) {
//	aui, err := adaptFirebaseUserInfoToAuthProviderUserInfo(ur.UserInfo)
//	if err != nil {
//		return nil, err
//	}
//	apui, err := adaptFirebaseUserInfoListToAuthProviderUserInfoList(ur.ProviderUserInfo)
//	if err != nil {
//		return nil, err
//	}
//	return &provider.User{
//		UserInfo:               aui,
//		CustomClaims:           ur.CustomClaims,
//		Disabled:               ur.Disabled,
//		EmailVerified:          ur.EmailVerified,
//		ProviderUserInfo:       apui,
//		TokensValidAfterMillis: ur.TokensValidAfterMillis,
//		UserMetadata:           adaptFirebaseUserMetadataToAuthProviderUserMetadata(ur.UserMetadata),
//		TenantID:               ur.TenantID,
//		MultiFactor: &provider.MultiFactorSettings{
//			EnrolledFactors: adaptFirebaseUserEnrolledFactorsToAuthProviderUserEnrolledFactors(ur.MultiFactor.EnrolledFactors),
//		},
//	}, nil
//}
