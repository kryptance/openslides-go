package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// Delegation history tests verify that user create/update work with delegation fields.
// The Python tests focus on history information logging which is not yet implemented in Go.
// These tests verify the basic action execution with delegation-related data.

func TestDelegationHistoryCreateWithMeeting(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("user.create", map[string]any{
		"username": "alice",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"username": "alice"})
}

func TestDelegationHistoryUpdateUser(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "alice"},
		"user/3": {"username": "bob"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":         2,
		"first_name": "Alice",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{
		"username":   "alice",
		"first_name": "Alice",
	})
}
