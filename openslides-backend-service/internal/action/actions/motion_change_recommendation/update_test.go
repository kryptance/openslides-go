package motion_change_recommendation

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/25": {
			"title":      "title_pheK0Ja3ai",
			"meeting_id": 1,
		},
		"motion_change_recommendation/111": {
			"line_from":  11,
			"line_to":    23,
			"text":       "text_LhmrbbwS",
			"motion_id":  25,
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_change_recommendation.update", map[string]any{
		"id":                111,
		"text":              "text_zzTWoMte",
		"rejected":          false,
		"internal":          true,
		"other_description": "other_description_IClpabuM",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_change_recommendation/111")
	if model["text"] != "text_zzTWoMte" {
		t.Errorf("expected text 'text_zzTWoMte', got %v", model["text"])
	}
	if model["rejected"] != false {
		t.Errorf("expected rejected false, got %v", model["rejected"])
	}
	if model["internal"] != true {
		t.Errorf("expected internal true, got %v", model["internal"])
	}
	if model["other_description"] != "other_description_IClpabuM" {
		t.Errorf("expected other_description, got %v", model["other_description"])
	}
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate model existence before updating.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/25": {
			"title":      "title_pheK0Ja3ai",
			"meeting_id": 1,
		},
		"motion_change_recommendation/111": {
			"line_from":  11,
			"line_to":    23,
			"text":       "text_LhmrbbwS",
			"motion_id":  25,
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_change_recommendation.update", map[string]any{
		"id": 112, "text": "text_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_change_recommendation/111", map[string]any{
		"text":      "text_LhmrbbwS",
		"line_from": 11,
		"line_to":   23,
		"motion_id": 25,
	})
}

func TestUpdateWithLineChanges(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"title":                     "Motion 1",
			"meeting_id":                1,
			"change_recommendation_ids": []any{1, 2, 3},
		},
		"motion_change_recommendation/1": {
			"meeting_id": 1, "motion_id": 1,
			"line_from": 1, "line_to": 2, "text": "Reco 1",
		},
		"motion_change_recommendation/2": {
			"meeting_id": 1, "motion_id": 1,
			"line_from": 4, "line_to": 6, "text": "Reco 2",
		},
		"motion_change_recommendation/3": {
			"meeting_id": 1, "motion_id": 1,
			"line_from": 8, "line_to": 10, "text": "Reco 3",
		},
	})
	// Update line_to for recommendation 1.
	resp, err := tc.Request("motion_change_recommendation.update", map[string]any{
		"id": 1, "line_to": 3,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_change_recommendation/1", map[string]any{
		"line_from": 1,
		"line_to":   3,
	})
}
