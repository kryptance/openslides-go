package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestPointOfOrderCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
	})

	resp, err := tc.Request("speaker.point_of_order", map[string]any{
		"list_of_speakers_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestPointOfOrderCreateWithNote(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
	})

	resp, err := tc.Request("speaker.point_of_order", map[string]any{
		"list_of_speakers_id": 1,
		"note":                "Important point",
	})
	tc.AssertSuccess(resp, err)
}

func TestPointOfOrderCreateWithCategory(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
		"point_of_order_category/1": {
			"meeting_id": 1,
			"text":       "Technical issue",
			"rank":       1,
		},
	})

	resp, err := tc.Request("speaker.point_of_order", map[string]any{
		"list_of_speakers_id":       1,
		"point_of_order_category_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestPointOfOrderCreateMissingListOfSpeakers(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.point_of_order", map[string]any{
		"note": "Some note",
	})
	tc.AssertErrorContains(err, "list_of_speakers_id")
}
