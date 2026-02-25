package motion_state

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/42": {
			"name":       "test_name_fjwnq8d8tje8",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_state.create", map[string]any{
		"name":        "test_Xcdfgee",
		"workflow_id": 42,
		"weight":      1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/42": {
			"name":       "test_name_fjwnq8d8tje8",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_state.create", map[string]any{
		"name":                   "test_Xcdfgee",
		"workflow_id":            42,
		"weight":                 1,
		"recommendation_label":   "recommend",
		"css_class":              "red",
		"allow_support":          true,
		"allow_create_poll":      true,
		"allow_submitter_edit":   true,
		"set_number":             true,
		"allow_motion_forwarding": true,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_state.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCreateMissingWorkflowId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_state.create", map[string]any{
		"name": "test_state",
	})
	tc.AssertError(resp, err)
}

func TestCreateMissingName(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_state.create", map[string]any{
		"workflow_id": 42,
	})
	tc.AssertError(resp, err)
}
