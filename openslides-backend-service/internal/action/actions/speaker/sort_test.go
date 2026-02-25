package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{1, 2, 3},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
		"speaker/1": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              1,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              2,
		},
		"speaker/3": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              3,
		},
	})

	resp, err := tc.Request("speaker.sort", map[string]any{
		"list_of_speakers_id": 1,
		"speaker_ids":         []any{3, 1, 2},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/3", map[string]any{"weight": 1})
	tc.AssertModelExists("speaker/1", map[string]any{"weight": 2})
	tc.AssertModelExists("speaker/2", map[string]any{"weight": 3})
}

func TestSortReverseOrder(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{1, 2},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
		"speaker/1": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              1,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              2,
		},
	})

	resp, err := tc.Request("speaker.sort", map[string]any{
		"list_of_speakers_id": 1,
		"speaker_ids":         []any{2, 1},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/2", map[string]any{"weight": 1})
	tc.AssertModelExists("speaker/1", map[string]any{"weight": 2})
}

func TestSortSingleSpeaker(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
			"speaker_ids":       []any{1},
		},
		"topic/1": {
			"meeting_id":          1,
			"title":               "Test Topic",
			"list_of_speakers_id": 1,
		},
		"speaker/1": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              5,
		},
	})

	resp, err := tc.Request("speaker.sort", map[string]any{
		"list_of_speakers_id": 1,
		"speaker_ids":         []any{1},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{"weight": 1})
}

func TestSortMissingSpeakerIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.sort", map[string]any{
		"list_of_speakers_id": 1,
	})
	tc.AssertErrorContains(err, "speaker_ids")
}

func TestSortMissingListOfSpeakersId(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.sort", map[string]any{
		"speaker_ids": []any{1, 2},
	})
	tc.AssertErrorContains(err, "list_of_speakers_id")
}

func TestSortEmptySpeakerIds(t *testing.T) {
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

	resp, err := tc.Request("speaker.sort", map[string]any{
		"list_of_speakers_id": 1,
		"speaker_ids":         []any{},
	})
	tc.AssertSuccess(resp, err)
}
