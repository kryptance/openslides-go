package list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestDeleteAllSpeakersCorrect(t *testing.T) {
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

	resp, err := tc.Request("list_of_speakers.delete_all_speakers", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/1")
	tc.AssertHasEvent(resp, event.TypeDelete, "speaker/2")
}

func TestDeleteAllSpeakersNoSpeakers(t *testing.T) {
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

	resp, err := tc.Request("list_of_speakers.delete_all_speakers", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 0)
}

func TestDeleteAllSpeakersNilSpeakerIds(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"list_of_speakers/1": {
			"meeting_id":        1,
			"content_object_id": "topic/1",
		},
	})

	resp, err := tc.Request("list_of_speakers.delete_all_speakers", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 0)
}

func TestDeleteAllSpeakersWrongID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("list_of_speakers.delete_all_speakers", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestDeleteAllSpeakersWithActiveSpeaker(t *testing.T) {
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
			"begin_time":          1000000,
			"end_time":            1000300,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"begin_time":          1000300,
		},
		"speaker/3": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              3,
		},
	})

	resp, err := tc.Request("list_of_speakers.delete_all_speakers", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 3)
}
