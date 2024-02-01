package workspace

import (
	"fmt"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"testing"
)

func invalidErrorMsg(v string) string {
	return fmt.Sprintf("invalid workspace domain: %s", v)
}

func TestNewDomain(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Subdomain
		wantErr validation.Errors
	}{
		{"success", "test", Subdomain{"test"}, nil},
		{"success", "valid123", Subdomain{"valid123"}, nil},
		{"success only number", "123", Subdomain{"123"}, nil},
		{"success kebab case", "example-domain", Subdomain{"example-domain"}, nil},
		{"success kebab case with number", "example-domain123", Subdomain{"example-domain123"}, nil},
		{"success kebab case with number", "example-domain-123", Subdomain{"example-domain-123"}, nil},
		{"fail", "", Subdomain{}, test.NewValidationErrors(InvalidWorkspaceSubdomain, invalidErrorMsg("")).Errs},
		{"fail", "not--valid-domain", Subdomain{}, test.NewValidationErrors(InvalidWorkspaceSubdomain, invalidErrorMsg("not--valid-domain")).Errs},
		{"fail", "-invalid", Subdomain{}, test.NewValidationErrors(InvalidWorkspaceSubdomain, invalidErrorMsg("-invalid")).Errs},
		{"fail", "invalid-", Subdomain{}, test.NewValidationErrors(InvalidWorkspaceSubdomain, invalidErrorMsg("invalid-")).Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSubdomain(tt.args)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewSubdomain() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewSubdomain() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
