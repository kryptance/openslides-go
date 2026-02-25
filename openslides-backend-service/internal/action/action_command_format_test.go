package action_test

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"

	// Ensure all needed action packages are registered.
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/group"
	_ "github.com/OpenSlides/openslides-backend-service/internal/action/actions/meeting"
)

// Tests migrated from test_action_command_format.py.
// These tests verify multi-action and multi-instance command processing.

func TestCreateMultipleInstancesInOneAction(t *testing.T) {
	// Corresponds to test_create_1_2_events: create 2 groups in 1 action call.
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.RequestMulti("group.create", []map[string]any{
		{"name": "group 1", "meeting_id": 1},
		{"name": "group 2", "meeting_id": 1},
	})
	tc.AssertSuccess(resp, err)
}

func TestUpdateMultipleInstancesInOneAction(t *testing.T) {
	// Corresponds to test_update_1_2_events: update 2 meetings in 1 action call.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"name":                         "name1",
			"committee_id":                 1,
			"welcome_title":                "t",
			"is_active_in_organization_id": 1,
			"language":                     "en",
		},
		"meeting/2": {
			"name":                         "name2",
			"committee_id":                 1,
			"welcome_title":                "t",
			"is_active_in_organization_id": 1,
			"language":                     "en",
		},
		"committee/1": {"name": "test_committee"},
	})
	resp, err := tc.RequestMulti("meeting.update", []map[string]any{
		{"id": 1, "name": "name1_updated"},
		{"id": 2, "name": "name2_updated"},
	})
	tc.AssertSuccess(resp, err)
	meeting1 := tc.GetModel("meeting/1")
	if meeting1["name"] != "name1_updated" {
		t.Errorf("expected meeting/1 name 'name1_updated', got %v", meeting1["name"])
	}
	meeting2 := tc.GetModel("meeting/2")
	if meeting2["name"] != "name2_updated" {
		t.Errorf("expected meeting/2 name 'name2_updated', got %v", meeting2["name"])
	}
}

func TestLookupNonExistentAction(t *testing.T) {
	// Corresponds to test_request_no_existing_action.
	_, err := action.Lookup("fuzzy_action_hamzaeNg4a")
	if err == nil {
		t.Fatal("expected error for non-existent action")
	}
}
