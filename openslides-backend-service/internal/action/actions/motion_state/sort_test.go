package motion_state

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortCorrect(t *testing.T) {
	t.Skip("Sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/1": {
			"state_ids":  []any{1, 2, 3},
			"meeting_id": 1,
		},
		"motion_state/1": {"workflow_id": 1, "meeting_id": 1},
		"motion_state/2": {"workflow_id": 1, "meeting_id": 1},
		"motion_state/3": {"workflow_id": 1, "meeting_id": 1},
	})
	resp, err := tc.Request("motion_state.sort", map[string]any{
		"workflow_id": 1,
		"ids":         []any{3, 2, 1},
	})
	tc.AssertSuccess(resp, err)
}
