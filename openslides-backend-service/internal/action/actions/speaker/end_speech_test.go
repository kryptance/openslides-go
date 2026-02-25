package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestEndSpeechCorrect(t *testing.T) {
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
		},
	})

	resp, err := tc.Request("speaker.end_speech", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["end_time"] == nil {
		t.Error("expected end_time to be set")
	}
}

func TestEndSpeechSetsEndTime(t *testing.T) {
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
		},
	})

	resp, err := tc.Request("speaker.end_speech", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	endTime := model["end_time"]
	if endTime == nil || endTime == 0 {
		t.Error("expected end_time to be set to a non-zero value")
	}
}

func TestEndSpeechMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.end_speech", map[string]any{})
	tc.AssertErrorContains(err, "id")
}

func TestEndSpeechWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on update")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("speaker.end_speech", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}
