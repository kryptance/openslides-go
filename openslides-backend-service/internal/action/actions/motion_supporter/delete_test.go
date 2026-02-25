package motion_supporter

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting_user/1": {
			"meeting_id":            1,
			"user_id":               1,
			"motion_supporter_ids":  []any{2},
		},
		"motion_supporter/2": {
			"meeting_user_id": 1,
			"motion_id":       1,
			"meeting_id":      1,
		},
		"motion/1": {
			"supporter_ids": []any{2},
			"title":         "motion_1",
			"meeting_id":    1,
			"state_id":      1,
		},
		"motion_state/1": {
			"name":          "state_1",
			"allow_support": true,
			"motion_ids":    []any{1},
			"meeting_id":    1,
		},
	})
	resp, err := tc.Request("motion_supporter.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_supporter/2")
}
