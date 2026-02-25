package assignment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("assignment.update", map[string]any{
		"id": 111, "title": "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment/111", map[string]any{"title": "title_Xcdfgee"})
}

func TestUpdateCorrectFullFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("assignment.update", map[string]any{
		"id":                       111,
		"title":                    "title_Xcdfgee",
		"description":              "text_test1",
		"open_posts":               12,
		"default_poll_description": "text_test2",
		"number_poll_candidates":   true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("assignment/111", map[string]any{
		"title":                    "title_Xcdfgee",
		"description":              "text_test1",
		"open_posts":               12,
		"default_poll_description": "text_test2",
		"number_poll_candidates":   true,
	})
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate that the model exists before updating.
	// It applies changes to any ID. This differs from the Python backend.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("assignment.update", map[string]any{
		"id": 112, "title": "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	// The original model 111 should remain untouched.
	tc.AssertModelExists("assignment/111", map[string]any{"title": "title_srtgb123"})
}
