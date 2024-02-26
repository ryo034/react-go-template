package middleware

import (
	"context"
	"fmt"

	fb "github.com/ryo034/react-go-template/apps/system/api/infrastructure/firebase"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/shared"
	"github.com/ryo034/react-go-template/apps/system/api/schema/openapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	firebase *fb.Firebase
	co       shared.ContextOperator
}

func NewSecMiddleware(firebase *fb.Firebase, co shared.ContextOperator) *Middleware {
	return &Middleware{firebase, co}
}

// HandleBearer run only when the authentication method is Bearer
func (m *Middleware) HandleBearer(ctx context.Context, operationName string, t openapi.Bearer) (context.Context, error) {
	return m.checkAndSetToken(ctx, t.GetToken())
}

// checkAndSetToken return JWT information in the key token in the argument context
func (m *Middleware) checkAndSetToken(ctx context.Context, t string) (context.Context, error) {
	token, err := m.firebase.Auth.VerifyIDToken(ctx, t)
	if err != nil || token == nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid JWT token"))
	}
	// Return JWT information in the key token in the argument context
	return m.co.SetClaim(m.co.SetUID(ctx, token.UID), token.Claims), nil
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
