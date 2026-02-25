package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestForgetPasswordNoEvents(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"url": "https://example.com"},
		"user/1":         {"email": "test@example.com"},
	})
	resp, err := tc.Request("user.forget_password", map[string]any{
		"email": "test@example.com",
	})
	tc.AssertSuccess(resp, err)
	// forget_password returns no events (side-effect is email sending)
	tc.AssertEventCount(resp, 0)
}

func TestForgetPasswordNoUserFound(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"url": "https://example.com"},
	})
	// Even if no user found, should return success (to avoid email enumeration)
	resp, err := tc.Request("user.forget_password", map[string]any{
		"email": "nonexistent@example.com",
	})
	tc.AssertSuccess(resp, err)
}

func TestForgetPasswordEmptyEmail(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.forget_password", map[string]any{
		"email": "",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "email")
}
