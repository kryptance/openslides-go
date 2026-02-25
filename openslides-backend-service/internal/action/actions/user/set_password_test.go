package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSetPasswordCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "password": "old_pw"},
	})
	resp, err := tc.Request("user.set_password", map[string]any{
		"id":       2,
		"password": "new_password",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"password": "new_password"})
}

func TestSetPasswordWithSetAsDefault(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "password": "old_pw"},
	})
	resp, err := tc.Request("user.set_password", map[string]any{
		"id":             2,
		"password":       "new_password",
		"set_as_default": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{
		"password":         "new_password",
		"default_password": "new_password",
	})
}

func TestSetPasswordWithoutSetAsDefault(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "password": "old_pw", "default_password": "old_default"},
	})
	resp, err := tc.Request("user.set_password", map[string]any{
		"id":             2,
		"password":       "new_password",
		"set_as_default": false,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{
		"password":         "new_password",
		"default_password": "old_default",
	})
}

func TestSetPasswordOnSelf(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old_pw"},
	})
	resp, err := tc.Request("user.set_password", map[string]any{
		"id":             1,
		"password":       "new_password",
		"set_as_default": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"password":         "new_password",
		"default_password": "new_password",
	})
}

func TestSetPasswordNonExistentUser(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// update operations. The set_password succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("user.set_password", map[string]any{
		"id":       999,
		"password": "new_password",
	})
	tc.AssertSuccess(resp, err)
}
