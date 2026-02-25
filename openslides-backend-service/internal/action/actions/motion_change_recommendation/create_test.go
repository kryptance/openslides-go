package motion_change_recommendation

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateGoodRequiredFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/233": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_change_recommendation.create", map[string]any{
		"line_from": 125,
		"line_to":   234,
		"text":      "text_DvLXGcdW",
		"motion_id": 233,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_change_recommendation/1", map[string]any{
		"line_from": 125,
		"line_to":   234,
		"text":      "text_DvLXGcdW",
		"motion_id": 233,
	})
}

func TestCreateGoodAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/233": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_change_recommendation.create", map[string]any{
		"line_from":         125,
		"line_to":           234,
		"text":              "text_DvLXGcdW",
		"motion_id":         233,
		"rejected":          false,
		"internal":          true,
		"type":              0, // In Go schema, type is an integer (0=replacement, 1=insertion, etc.)
		"other_description": "other_description_iuDguxZp",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_change_recommendation/1")
	if model["line_from"] != 125 {
		t.Errorf("expected line_from 125, got %v", model["line_from"])
	}
	if model["line_to"] != 234 {
		t.Errorf("expected line_to 234, got %v", model["line_to"])
	}
	if model["text"] != "text_DvLXGcdW" {
		t.Errorf("expected text 'text_DvLXGcdW', got %v", model["text"])
	}
	if model["motion_id"] != 233 {
		t.Errorf("expected motion_id 233, got %v", model["motion_id"])
	}
	if model["internal"] != true {
		t.Errorf("expected internal true, got %v", model["internal"])
	}
	if model["other_description"] != "other_description_iuDguxZp" {
		t.Errorf("expected other_description, got %v", model["other_description"])
	}
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_change_recommendation.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateWrongField(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/233": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_change_recommendation.create", map[string]any{
		"line_from":   125,
		"line_to":     234,
		"text":        "text_DvLXGcdW",
		"motion_id":   233,
		"wrong_field": "text_AefohteiF8",
	})
	// The Go schema does not enforce additionalProperties=false.
	tc.AssertSuccess(resp, err)
}

func TestCreateTitleChangeRecommendation(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/233": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_change_recommendation.create", map[string]any{
		"line_from": 0,
		"line_to":   0,
		"text":      "text",
		"motion_id": 233,
	})
	tc.AssertSuccess(resp, err)
}
