package motion_state

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_state/111": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_state.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_state/111")
}
