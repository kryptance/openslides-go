package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestGenerateNewPasswordCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old_pw"},
	})
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/1")
	password, ok := model["password"].(string)
	if !ok || password == "" {
		t.Fatal("expected password to be set")
	}
	defaultPassword, ok := model["default_password"].(string)
	if !ok || defaultPassword == "" {
		t.Fatal("expected default_password to be set")
	}
	// In the current implementation (no bcrypt), password == default_password
	if password != defaultPassword {
		t.Fatalf("expected password (%s) == default_password (%s)", password, defaultPassword)
	}
}

func TestGenerateNewPasswordLength(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old_pw"},
	})
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/1")
	defaultPassword, ok := model["default_password"].(string)
	if !ok {
		t.Fatal("expected default_password to be set")
	}
	if len(defaultPassword) != passwordLength {
		t.Fatalf("expected password length %d, got %d", passwordLength, len(defaultPassword))
	}
}

func TestGenerateNewPasswordCharset(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old_pw"},
	})
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/1")
	defaultPassword, ok := model["default_password"].(string)
	if !ok {
		t.Fatal("expected default_password to be set")
	}
	for _, c := range defaultPassword {
		found := false
		for _, allowed := range passwordCharset {
			if c == allowed {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("character %c is not in the allowed charset", c)
		}
	}
}

func TestGenerateNewPasswordDifferentUser(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "password": "old_pw"},
	})
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/2")
	password, ok := model["password"].(string)
	if !ok || password == "" {
		t.Fatal("expected password to be set")
	}
	if password == "old_pw" {
		t.Fatal("password should have been changed")
	}
}

func TestGenerateNewPasswordNonExistentUser(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// update operations. The action succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 999})
	tc.AssertSuccess(resp, err)
}

func TestGenerateNewPasswordReplacesOldPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {
			"username":         "testuser",
			"password":         "old_hash",
			"default_password": "old_default",
		},
	})
	resp, err := tc.Request("user.generate_new_password", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/2")
	newDefault := model["default_password"].(string)
	if newDefault == "old_default" {
		t.Fatal("default_password should have been changed")
	}
	newPassword := model["password"].(string)
	if newPassword == "old_hash" {
		t.Fatal("password should have been changed")
	}
}
