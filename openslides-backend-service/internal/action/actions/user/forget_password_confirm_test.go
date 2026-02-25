package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestForgetPasswordConfirmCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"password": "old", "email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password_confirm", map[string]any{
		"user_id":      1,
		"token":        "valid_token",
		"new_password": "new_password",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeUpdate, "user/1")
}

func TestForgetPasswordConfirmEmptyToken(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"password": "old", "email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password_confirm", map[string]any{
		"user_id":      1,
		"token":        "",
		"new_password": "new_password",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "token")
}

func TestForgetPasswordConfirmEmptyPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"password": "old", "email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password_confirm", map[string]any{
		"user_id":      1,
		"token":        "valid_token",
		"new_password": "",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "new_password")
}

func TestForgetPasswordConfirmSetsPasswordField(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"password": "old", "email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password_confirm", map[string]any{
		"user_id":      2,
		"token":        "valid_token",
		"new_password": "brand_new_pw",
	})
	tc.AssertSuccess(resp, err)
	if resp == nil || len(resp.Events) == 0 {
		t.Fatal("expected events")
	}
	e := resp.Events[0]
	if e.Fields["password"] != "brand_new_pw" {
		t.Fatalf("expected password field to be brand_new_pw, got %v", e.Fields["password"])
	}
}

func TestForgetPasswordConfirmEventCount(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"password": "old", "email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password_confirm", map[string]any{
		"user_id":      1,
		"token":        "valid_token",
		"new_password": "new_password",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 1)
}
