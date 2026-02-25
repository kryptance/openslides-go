package motion_workflow

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/42": {
			"name":                        "test_name_fsdksjdfhdsfssdf",
			"is_active_in_organization_id": 1,
			"committee_id":                1,
		},
	})
	resp, err := tc.Request("motion_workflow.create", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 42,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_workflow.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCreateMissingName(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_workflow.create", map[string]any{
		"meeting_id": 42,
	})
	tc.AssertError(resp, err)
}

func TestCreateMissingMeetingId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_workflow.create", map[string]any{
		"name": "test_Xcdfgee",
	})
	tc.AssertError(resp, err)
}
