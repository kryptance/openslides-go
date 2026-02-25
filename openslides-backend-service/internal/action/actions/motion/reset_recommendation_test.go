package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestResetRecommendationCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_state/77": {
			"meeting_id":                 1,
			"name":                       "test1",
			"motion_recommendation_ids":  []any{22},
		},
		"motion/22": {
			"meeting_id":        1,
			"title":             "test1",
			"recommendation_id": 77,
		},
	})
	resp, err := tc.Request("motion.reset_recommendation", map[string]any{"id": 22})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion/22")
	if model["recommendation_id"] != nil {
		t.Errorf("expected recommendation_id to be nil, got %v", model["recommendation_id"])
	}
}

func TestResetRecommendationCorrectEmptyRecommendation(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_state/77": {
			"meeting_id": 1,
			"name":       "test1",
		},
		"motion/22": {
			"meeting_id": 1,
			"title":      "test1",
		},
	})
	resp, err := tc.Request("motion.reset_recommendation", map[string]any{"id": 22})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion/22")
	if model["recommendation_id"] != nil {
		t.Errorf("expected recommendation_id to be nil, got %v", model["recommendation_id"])
	}
}
