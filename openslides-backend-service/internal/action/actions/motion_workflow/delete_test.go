package motion_workflow

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/111": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_workflow.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_workflow/111")
}

func TestDeleteWithStates(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/2": {
			"meeting_id": 1,
			"state_ids":  []any{3},
		},
		"motion_state/3": {
			"workflow_id": 2,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion_workflow.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_workflow/2")
}
