package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSpeakCorrect(t *testing.T) {
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
			"weight":              1,
		},
	})

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["begin_time"] == nil || model["begin_time"] == 0 {
		t.Error("expected begin_time to be set")
	}
}

func TestSpeakSetsBeginTime(t *testing.T) {
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
		},
	})

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["begin_time"] == nil {
		t.Error("expected begin_time to be set after speaking")
	}
}

func TestSpeakEndsCurrentSpeaker(t *testing.T) {
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
			"begin_time":          1000000,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              2,
		},
	})

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 2,
	})
	tc.AssertSuccess(resp, err)

	// The previous speaker should have end_time set.
	model1 := tc.GetModel("speaker/1")
	if model1["end_time"] == nil {
		t.Error("expected speaker/1 end_time to be set after new speaker begins")
	}

	// The new speaker should have begin_time set.
	model2 := tc.GetModel("speaker/2")
	if model2["begin_time"] == nil {
		t.Error("expected speaker/2 begin_time to be set")
	}
}

func TestSpeakNoCurrentSpeaker(t *testing.T) {
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
			"weight":              1,
		},
	})

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("speaker/1")
	if model["begin_time"] == nil {
		t.Error("expected begin_time to be set")
	}
}

func TestSpeakWrongID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestSpeakDoesNotEndAlreadyFinishedSpeaker(t *testing.T) {
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
			"begin_time":          1000000,
			"end_time":            1000300,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"weight":              2,
		},
	})

	resp, err := tc.Request("speaker.speak", map[string]any{
		"id": 2,
	})
	tc.AssertSuccess(resp, err)

	// Already finished speaker should keep its original end_time.
	model1 := tc.GetModel("speaker/1")
	if model1["end_time"] != 1000300 {
		t.Errorf("expected speaker/1 end_time to remain 1000300, got %v", model1["end_time"])
	}
}
