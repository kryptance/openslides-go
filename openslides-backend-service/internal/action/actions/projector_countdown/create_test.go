package projector_countdown

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"projector_countdown_default_time": 11,
		},
	})

	resp, err := tc.Request("projector_countdown.create", map[string]any{
		"meeting_id":   1,
		"title":        "test",
		"description":  "good description",
		"default_time": 30,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_countdown/1")
	if model["title"] != "test" {
		t.Errorf("expected title 'test', got %v", model["title"])
	}
	if model["meeting_id"] != 1 {
		t.Errorf("expected meeting_id 1, got %v", model["meeting_id"])
	}
	if model["description"] != "good description" {
		t.Errorf("expected description 'good description', got %v", model["description"])
	}
}

func TestCreateWithCountdownTime(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("projector_countdown.create", map[string]any{
		"meeting_id":     1,
		"title":          "test2",
		"description":    "good description",
		"default_time":   30,
		"countdown_time": 20,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_countdown/1")
	if model["title"] != "test2" {
		t.Errorf("expected title 'test2', got %v", model["title"])
	}
}
