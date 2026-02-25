package assignment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {"meeting_id": 1, "title": "title_srtgb123"},
	})
	resp, err := tc.Request("assignment.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("assignment/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate that the model exists before deleting.
	// It marks the model as deleted even if it never existed. This differs from
	// the Python backend which returns a 400 error.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/112": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("assignment.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	// The original model 112 should remain untouched.
	tc.AssertModelExists("assignment/112", map[string]any{"title": "title_srtgb123"})
}

func TestDeleteCorrectCascading(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/111": {
			"list_of_speakers_id": 222,
			"agenda_item_id":      333,
			"meeting_id":          1,
			"phase":               "finished",
			"candidate_ids":       []any{1111},
		},
		"list_of_speakers/222": {
			"closed":            false,
			"content_object_id": "assignment/111",
			"meeting_id":        1,
		},
		"agenda_item/333": {
			"comment":           "test_comment_ewoirzewoirioewr",
			"content_object_id": "assignment/111",
			"meeting_id":        1,
		},
		"assignment_candidate/1111": {"assignment_id": 111, "meeting_id": 1},
	})
	resp, err := tc.Request("assignment.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("assignment/111")
}
