package projector_countdown

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"projector_countdown/2": {
			"meeting_id":     1,
			"title":          "test",
			"description":    "blablabla",
			"default_time":   60,
			"countdown_time": 60,
		},
	})

	resp, err := tc.Request("projector_countdown.update", map[string]any{
		"id":             2,
		"title":          "new_title",
		"description":    "good bla",
		"default_time":   30,
		"countdown_time": 20,
		"running":        true,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_countdown/2")
	if model["title"] != "new_title" {
		t.Errorf("expected title 'new_title', got %v", model["title"])
	}
	if model["description"] != "good bla" {
		t.Errorf("expected description 'good bla', got %v", model["description"])
	}
}

func TestUpdateSameTitleInSameID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"projector_countdown/2": {
			"meeting_id":     1,
			"title":          "test",
			"description":    "blablabla",
			"default_time":   60,
			"countdown_time": 60,
		},
	})

	resp, err := tc.Request("projector_countdown.update", map[string]any{
		"id":    2,
		"title": "test",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_countdown/2")
	if model["title"] != "test" {
		t.Errorf("expected title 'test', got %v", model["title"])
	}
}
