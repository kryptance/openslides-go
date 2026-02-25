package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestResetPasswordToDefault(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {
			"username":         "username_srtgb123",
			"default_password": "pw_quSEYapV",
			"password":         "some_hashed_pw",
		},
	})
	resp, err := tc.Request("user.reset_password_to_default", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/111", map[string]any{
		"password":         "pw_quSEYapV",
		"default_password": "pw_quSEYapV",
	})
}

func TestResetPasswordToDefaultSetsPasswordToDefault(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {
			"username":         "testuser",
			"default_password": "my_default_pass",
			"password":         "current_hash",
		},
	})
	resp, err := tc.Request("user.reset_password_to_default", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/2")
	password := model["password"].(string)
	defaultPassword := model["default_password"].(string)
	if password != defaultPassword {
		t.Fatalf("expected password (%s) == default_password (%s)", password, defaultPassword)
	}
}

func TestResetPasswordToDefaultNoDefaultPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {
			"username": "username_srtgb123",
			"password": "some_hash",
		},
	})
	resp, err := tc.Request("user.reset_password_to_default", map[string]any{"id": 111})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "default_password")
}

func TestResetPasswordToDefaultNonExistentUser(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.reset_password_to_default", map[string]any{"id": 999})
	tc.AssertError(resp, err)
}

func TestResetPasswordToDefaultSelfUser(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {
			"username":         "admin",
			"default_password": "admin_default",
			"password":         "current",
		},
	})
	resp, err := tc.Request("user.reset_password_to_default", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"password": "admin_default",
	})
}
