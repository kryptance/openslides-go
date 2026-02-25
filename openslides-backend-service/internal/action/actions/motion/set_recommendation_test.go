package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSetRecommendationCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_workflow/34": {
			"meeting_id": 1,
		},
		"motion_state/66": {
			"meeting_id":  1,
			"motion_ids":  []any{22},
			"workflow_id": 34,
		},
		"motion_state/77": {
			"meeting_id":           1,
			"workflow_id":          34,
			"recommendation_label": "blablabal",
		},
		"motion/22": {
			"meeting_id": 1,
			"state_id":   66,
		},
	})
	resp, err := tc.Request("motion.set_recommendation", map[string]any{
		"id":                22,
		"recommendation_id": 77,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion/22", map[string]any{"recommendation_id": 77})
}
