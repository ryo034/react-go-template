package user

import (
	"testing"

	"github.com/ryo034/react-go-template/apps/system/api/util/test"

	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
)

func TestAdapter_Adapt(t *testing.T) {
	validUUID := uuid.New()
	validEmail := "test@example.com"
	validPhoneNumber := "+819012341234"
	defaultTime := test.GetDefaultTime()

	anID, _ := uuid.NewV7()
	aeID, _ := uuid.NewV7()
	apnID, _ := uuid.NewV7()

	systemAccount := models.Account{
		AccountID: validUUID,
		Name: &models.AccountLatestName{
			AccountNameID: anID,
			AccountID:     validUUID,
			AccountName: &models.AccountName{
				AccountNameID: anID,
				AccountID:     validUUID,
				Name:          "John Doe",
				CreatedAt:     defaultTime,
			},
		},
		Email: &models.AccountLatestEmail{
			AccountEmailID: aeID,
			AccountID:      validUUID,
			AccountEmail: &models.AccountEmail{
				AccountEmailID: aeID,
				AccountID:      validUUID,
				Email:          validEmail,
				CreatedAt:      defaultTime,
			},
		},
		PhoneNumber: &models.AccountLatestPhoneNumber{
			AccountPhoneNumberID: apnID,
			AccountID:            validUUID,
			AccountPhoneNumber: &models.AccountPhoneNumber{
				AccountPhoneNumberID: apnID,
				AccountID:            validUUID,
				PhoneNumber:          validPhoneNumber,
				CreatedAt:            defaultTime,
			},
		},
	}

	adap := NewAdapter()

	t.Run("Valid conversion", func(t *testing.T) {
		u, err := adap.AdaptTmp(&systemAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})

	t.Run("PhoneNumber is nil", func(t *testing.T) {
		invalidAccount := systemAccount
		invalidAccount.PhoneNumber = nil
		u, err := adap.AdaptTmp(&invalidAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})

	t.Run("Name is empty", func(t *testing.T) {
		invalidAccount := systemAccount
		invalidAccount.Name = nil
		u, err := adap.AdaptTmp(&invalidAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})
}
