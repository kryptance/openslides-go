package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSetPasswordSelfCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old_password"},
	})
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "old_password",
		"new_password": "new_password",
	})
	tc.AssertSuccess(resp, err)
	// The action sets password = new_password (placeholder without bcrypt)
	tc.AssertModelExists("user/1", map[string]any{"password": "new_password"})
}

func TestSetPasswordSelfSetsIdToCurrentUser(t *testing.T) {
	tc := testutil.New(t)
	tc.UserID = 5
	tc.SetModels(map[string]map[string]any{
		"user/5": {"username": "user5", "password": "old"},
	})
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "old",
		"new_password": "new",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/5", map[string]any{"password": "new"})
}

func TestSetPasswordSelfAnonymous(t *testing.T) {
	tc := testutil.New(t)
	tc.UserID = 0 // anonymous
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "old",
		"new_password": "new",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "nonymous")
}

func TestSetPasswordSelfEmptyNewPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old"},
	})
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "old",
		"new_password": "",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "new_password")
}

func TestSetPasswordSelfEmptyOldPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old"},
	})
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "",
		"new_password": "new",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "old_password")
}

func TestSetPasswordSelfRemovesActionFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin", "password": "old"},
	})
	resp, err := tc.Request("user.set_password_self", map[string]any{
		"old_password": "old",
		"new_password": "new",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("user/1")
	// old_password and new_password should not be stored on the model
	if _, ok := model["old_password"]; ok {
		t.Error("old_password should not be stored on the user model")
	}
	if _, ok := model["new_password"]; ok {
		t.Error("new_password should not be stored on the user model")
	}
}
