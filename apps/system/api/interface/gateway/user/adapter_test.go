package user

import (
	"github.com/google/uuid"
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/database/bun/models"
	"testing"
)

func TestAdapter_Adapt(t *testing.T) {
	validUUID := uuid.New()
	validEmail := "test@example.com"
	validName := "John Doe"
	validPhoneNumber := "09000000000"

	systemAccount := models.SystemAccount{
		SystemAccountID: validUUID,
		Profile: &models.SystemAccountProfile{
			Email: validEmail,
			Name:  validName,
		},
		PhoneNumber: &models.SystemAccountPhoneNumber{
			PhoneNumber: validPhoneNumber,
		},
	}

	adapter := NewAdapter()

	t.Run("Valid conversion", func(t *testing.T) {
		u, err := adapter.Adapt(&systemAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})

	t.Run("PhoneNumber is nil", func(t *testing.T) {
		invalidSystemAccount := systemAccount
		invalidSystemAccount.PhoneNumber = nil
		u, err := adapter.Adapt(&invalidSystemAccount)
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
		u, err := adapter.Adapt(&invalidSystemAccount)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if u == nil {
			t.Fatal("Expected a user object, got nil")
		}
	})
}
