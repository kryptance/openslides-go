package topic

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {"title": "test", "meeting_id": 1},
	})

	resp, err := tc.Request("topic.update", map[string]any{
		"id":    1,
		"title": "test2",
		"text":  "text",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("topic/1")
	if model["title"] != "test2" {
		t.Errorf("expected title 'test2', got %v", model["title"])
	}
	if model["text"] != "text" {
		t.Errorf("expected text 'text', got %v", model["text"])
	}
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("topic.update", map[string]any{
		"title": "test2",
	})
	tc.AssertErrorContains(err, "id")
}

func TestUpdateTitleOnly(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/1": {"title": "original", "text": "old text", "meeting_id": 1},
	})

	resp, err := tc.Request("topic.update", map[string]any{
		"id":    1,
		"title": "updated",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("topic/1")
	if model["title"] != "updated" {
		t.Errorf("expected title 'updated', got %v", model["title"])
	}
}
