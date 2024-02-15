package workspace

import (
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/validation"
	"github.com/ryo034/react-go-template/apps/system/api/util/test"
	"reflect"
	"strings"
	"testing"
)

func TestNewName_OK(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Name
		wantErr bool
	}{
		{"success", "test", Name{"test"}, false},
		{"hyphen", "test-name", Name{"test-name"}, false},
		{"japanese", "あいうえお", Name{"あいうえお"}, false},
		{"japanese kana", "カタカナ", Name{"カタカナ"}, false},
		{"kanji", "漢字", Name{"漢字"}, false},
		{"trim", " test ", Name{"test"}, false},
		{"max 255", strings.Repeat("a", 255), Name{strings.Repeat("a", 255)}, false},
		{"contains number", "test123", Name{"test123"}, false},
		{"contains number with hyphen", "test-123", Name{"test-123"}, false},
		{"under bar", "test_123", Name{"test_123"}, false},
		{"zenkaku space", "株式会社　テスト", Name{"株式会社　テスト"}, false},
		{"contains number with space", "test 123", Name{"test 123"}, false},
		{"has space", "test name", Name{"test name"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewName_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Name
		wantErr validation.Errors
	}{
		{"empty", "", Name{}, test.NewValidationErrors(InvalidWorkspaceName, "").Errs},
		{"only space", " ", Name{}, test.NewValidationErrors(InvalidWorkspaceName, " ").Errs},
		{"over max 255", strings.Repeat("a", 256), Name{}, test.NewValidationErrors(InvalidWorkspaceName, strings.Repeat("a", 256)).Errs},
		{"over max 255 japanese", strings.Repeat("あ", 256), Name{}, test.NewValidationErrors(InvalidWorkspaceName, strings.Repeat("あ", 256)).Errs},
		{"invalid char", "test!", Name{}, test.NewValidationErrors(InvalidWorkspaceName, "test!").Errs},
		{"emoji", "🍣", Name{}, test.NewValidationErrors(InvalidWorkspaceName, "🍣").Errs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.value)
			if err == nil {
				if tt.wantErr != nil {
					t.Errorf("NewName() got = %v, want %v", got, tt.want)
					return
				}
			} else {
				if !test.ValidationErrorEquals(err.(validation.Errors), tt.wantErr) {
					t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
