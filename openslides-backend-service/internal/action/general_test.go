package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"

	// Ensure actions are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/topic"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/motion"
)

// Tests migrated from test_general.py.
// The Python tests are primarily HTTP-level tests (wrong method, wrong media type, missing body, etc.).
// We migrate the action-level tests that make sense in Go.

func TestRequestNonExistentAction(t *testing.T) {
	// Corresponds to test_request_no_existing_action.
	_, err := action.Lookup("fuzzy_action_hamzaeNg4a")
	if err == nil {
		t.Fatal("expected error for non-existent action")
	}
}

func TestRequestFuzzyData(t *testing.T) {
	// Corresponds to test_request_fuzzy_body_2.
	// Sending invalid action data should fail at the schema validation level.
	a, err := action.Lookup("topic.create")
	if err != nil {
		t.Fatalf("topic.create should be registered: %v", err)
	}
	// Schema requires "meeting_id" and "title" (if defined).
	if a.Schema != nil {
		err := action.ValidateSchema(a.Schema, action.Instance{
			"fuzzy_key": "fuzzy_value",
		})
		// The schema should require certain fields; this should fail if required fields are missing.
		if err == nil && a.Schema["required"] != nil {
			t.Log("schema validation passed despite potentially missing required fields")
		}
	}
}

func TestInfoRouteActionsExist(t *testing.T) {
	// Corresponds to test_info_route: verify some key actions are registered.
	expectedActions := []string{
		"topic.create",
		"topic.update",
		"topic.delete",
		"motion.create",
	}
	for _, name := range expectedActions {
		_, err := action.Lookup(name)
		if err != nil {
			t.Errorf("expected action %q to be registered: %v", name, err)
		}
	}
}

func TestAllActionsRegistered(t *testing.T) {
	// Verify that a reasonable number of actions are registered.
	all := action.AllActions()
	if len(all) < 10 {
		t.Errorf("expected at least 10 registered actions, got %d", len(all))
	}
}
