package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"meeting_id":          1,
			"title":               "title_srtgb123",
			"number":              "123",
			"text":                "<i>test</i>",
			"reason":              "<b>test2</b>",
			"modified_final_version": "blablabla",
		},
	})
	resp, err := tc.Request("motion.update", map[string]any{
		"id":                    111,
		"title":                 "title_bDFsWtKL",
		"number":                "124",
		"text":                  "text_eNPkDVuq",
		"reason":                "reason_ukWqADfE",
		"modified_final_version": "mfv_ilVvBsUi",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion/111")
	if model["title"] != "title_bDFsWtKL" {
		t.Errorf("expected title 'title_bDFsWtKL', got %v", model["title"])
	}
	if model["number"] != "124" {
		t.Errorf("expected number '124', got %v", model["number"])
	}
	if model["text"] != "text_eNPkDVuq" {
		t.Errorf("expected text 'text_eNPkDVuq', got %v", model["text"])
	}
	if model["reason"] != "reason_ukWqADfE" {
		t.Errorf("expected reason 'reason_ukWqADfE', got %v", model["reason"])
	}
}

func TestUpdateTitleOnly(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"meeting_id": 1,
			"title":      "old_title",
			"text":       "old_text",
		},
	})
	resp, err := tc.Request("motion.update", map[string]any{
		"id":    111,
		"title": "new_title",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion/111", map[string]any{"title": "new_title"})
}

func TestUpdateRecommendationExtension(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"meeting_id": 1,
			"title":      "title",
		},
	})
	resp, err := tc.Request("motion.update", map[string]any{
		"id":                       111,
		"recommendation_extension": "test extension",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion/111", map[string]any{
		"recommendation_extension": "test extension",
	})
}
