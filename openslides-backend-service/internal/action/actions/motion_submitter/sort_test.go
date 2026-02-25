package motion_submitter

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortCorrect(t *testing.T) {
	t.Skip("Sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/222": {"meeting_id": 1},
		"motion_submitter/31": {"motion_id": 222, "meeting_id": 1},
		"motion_submitter/32": {"motion_id": 222, "meeting_id": 1},
	})
	resp, err := tc.Request("motion_submitter.sort", map[string]any{
		"motion_id": 222,
		"ids":       []any{32, 31},
	})
	tc.AssertSuccess(resp, err)
}
