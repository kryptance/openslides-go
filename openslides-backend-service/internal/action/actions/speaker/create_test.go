package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_user_id":     1,
	})
}

func TestCreateWithSpeechState(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"speech_state":        "contribution",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "contribution",
	})
}

func TestCreateWithSpeechStatePro(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"speech_state":        "pro",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "pro",
	})
}

func TestCreateWithSpeechStateContra(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"speech_state":        "contra",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "contra",
	})
}

func TestCreateWithSpeechStateIntervention(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"speech_state":        "intervention",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "intervention",
	})
}

func TestCreateWithPointOfOrder(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"point_of_order":      true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"point_of_order": true,
	})
}

func TestCreateWithNote(t *testing.T) {
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

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
		"note":                "A speaker note",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"note": "A speaker note",
	})
}

func TestCreateMissingListOfSpeakers(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.create", map[string]any{
		"meeting_id":      1,
		"meeting_user_id": 1,
	})
	tc.AssertErrorContains(err, "list_of_speakers_id")
}

func TestCreateMissingMeetingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_user_id":     1,
	})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestCreateMultipleSpeakers(t *testing.T) {
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
		"user/2": {
			"username":         "user2",
			"is_active":       true,
			"organization_id": 1,
			"meeting_user_ids": []any{2},
		},
		"meeting_user/2": {
			"user_id":    2,
			"meeting_id": 1,
			"group_ids":  []any{1},
		},
	})

	resp, err := tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     1,
	})
	tc.AssertSuccess(resp, err)

	resp, err = tc.Request("speaker.create", map[string]any{
		"list_of_speakers_id": 1,
		"meeting_id":          1,
		"meeting_user_id":     2,
	})
	tc.AssertSuccess(resp, err)
}
