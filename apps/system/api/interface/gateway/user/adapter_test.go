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
	validName := "John Doe"
	validPhoneNumber := "09000000000"
	defaultTime := test.GetDefaultTime()

	systemAccount := models.SystemAccount{
		SystemAccountID: validUUID,
		Profile: &models.SystemAccountProfile{
			Name: validName,
		},
		Emails: []*models.SystemAccountEmail{
			{
				SystemAccountID: validUUID,
				Email:           validEmail,
				CreatedAt:       defaultTime,
			},
		},
		PhoneNumbers: []*models.SystemAccountPhoneNumber{
			{
				SystemAccountID: validUUID,
				PhoneNumber:     validPhoneNumber,
				CreatedAt:       defaultTime,
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
		invalidSystemAccount := systemAccount
		invalidSystemAccount.PhoneNumbers = nil
		u, err := adap.AdaptTmp(&invalidSystemAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})

	t.Run("Name is empty", func(t *testing.T) {
		invalidSystemAccount := systemAccount
		invalidSystemAccount.Profile.Name = ""
		u, err := adap.AdaptTmp(&invalidSystemAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})
}
