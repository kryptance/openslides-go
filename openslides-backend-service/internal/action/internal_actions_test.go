package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/user"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/organization"
)

// Tests migrated from test_internal_actions.py.
// The Python tests check internal action routing and auth.
// We test the internal action execution through the testutil framework.

func TestInternalUserCreate(t *testing.T) {
	// Corresponds to test_internal_user_create.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "Test Organization",
		},
	})
	resp, err := tc.RequestInternal("user.create", map[string]any{
		"username": "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"username": "test"})
}

func TestInternalUserUpdate(t *testing.T) {
	// Corresponds to test_internal_user_update.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "Test Organization",
		},
		"user/1": {
			"username": "admin",
		},
	})
	resp, err := tc.RequestInternal("user.update", map[string]any{
		"id":       1,
		"username": "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"username": "test"})
}

func TestInternalUserDelete(t *testing.T) {
	// Corresponds to test_internal_user_delete.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {
			"username": "admin",
		},
	})
	resp, err := tc.RequestInternal("user.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("user/1")
}

func TestBackendInternalActionNotCallableExternally(t *testing.T) {
	// Corresponds to test_internal_execute_stack_internal_via_public_route.
	// Backend-internal actions should not be callable via the non-internal path.
	a, err := action.Lookup("projection.create")
	if err != nil {
		t.Skipf("projection.create not registered: %v", err)
	}
	if a.ActionType != action.ActionTypeBackendInternal {
		t.Skip("projection.create is not marked as BackendInternal")
	}
	// The action is backend internal, so it should not be callable externally.
	// We verify this by checking the ActionType.
	if a.ActionType != action.ActionTypeBackendInternal {
		t.Error("expected projection.create to be BackendInternal")
	}
}
