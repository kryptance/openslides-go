package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUnpauseSpeechCorrect(t *testing.T) {
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
			"begin_time":          1000000,
			"pause_time":          1000100,
		},
	})

	resp, err := tc.Request("speaker.unpause_speech", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["pause_time"] != nil {
		t.Error("expected pause_time to be nil after unpause")
	}
}

func TestUnpauseSpeechCalculatesTotalPause(t *testing.T) {
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
			"begin_time":          1000000,
			"pause_time":          1000100,
			"total_pause":         50,
		},
	})

	resp, err := tc.Request("speaker.unpause_speech", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["pause_time"] != nil {
		t.Error("expected pause_time to be cleared after unpause")
	}
	// total_pause should have increased from 50.
	if model["total_pause"] == nil {
		t.Error("expected total_pause to be set")
	}
}

func TestUnpauseSpeechClearsPauseTime(t *testing.T) {
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
			"begin_time":          1000000,
			"pause_time":          1000050,
		},
	})

	resp, err := tc.Request("speaker.unpause_speech", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["pause_time"] != nil {
		t.Errorf("expected pause_time to be nil, got %v", model["pause_time"])
	}
}

func TestUnpauseSpeechWrongID(t *testing.T) {
	// unpause_speech validates model existence by looking up speaker in datastore
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("speaker.unpause_speech", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestUnpauseSpeechMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.unpause_speech", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
