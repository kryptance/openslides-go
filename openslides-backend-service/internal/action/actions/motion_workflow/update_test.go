package motion_workflow

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/111": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_workflow.update", map[string]any{
		"id":   111,
		"name": "name_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_workflow/111", map[string]any{"name": "name_Xcdfgee"})
}
