package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestFollowRecommendationCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_state/76": {
			"meeting_id":     1,
			"name":           "test0",
			"motion_ids":     []any{},
			"next_state_ids": []any{77},
		},
		"motion_state/77": {
			"meeting_id":                1,
			"name":                      "test1",
			"motion_ids":                []any{22},
			"first_state_of_workflow_id": 76,
			"previous_state_ids":        []any{76},
		},
		"motion/22": {
			"meeting_id":        1,
			"title":             "test1",
			"state_id":          77,
			"recommendation_id": 76,
		},
	})
	resp, err := tc.Request("motion.follow_recommendation", map[string]any{"id": 22})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion/22", map[string]any{"state_id": 76})
}

func TestFollowRecommendationMissingRecommendationId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_state/76": {
			"meeting_id":     1,
			"name":           "test0",
			"next_state_ids": []any{77},
		},
		"motion_state/77": {
			"meeting_id":                1,
			"name":                      "test1",
			"motion_ids":                []any{22},
			"first_state_of_workflow_id": 76,
			"previous_state_ids":        []any{76},
		},
		"motion/22": {
			"meeting_id": 1,
			"title":      "test1",
			"state_id":   77,
		},
	})
	resp, err := tc.Request("motion.follow_recommendation", map[string]any{"id": 22})
	// The Go implementation returns an error when recommendation_id is not set
	tc.AssertError(resp, err)
}
