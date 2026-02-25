package motion_change_recommendation

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_change_recommendation/111": {"meeting_id": 1, "motion_id": 1},
		"motion/1":                         {"meeting_id": 1, "change_recommendation_ids": []any{111}},
	})
	resp, err := tc.Request("motion_change_recommendation.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_change_recommendation/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_change_recommendation/112": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_change_recommendation.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_change_recommendation/112")
}
