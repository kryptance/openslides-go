package motion_working_group_speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/12": {
			"meeting_id": 1,
			"title":      "test2",
		},
		"motion_working_group_speaker/111": {
			"weight":     10,
			"motion_id":  12,
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_working_group_speaker.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_working_group_speaker/111")
}
