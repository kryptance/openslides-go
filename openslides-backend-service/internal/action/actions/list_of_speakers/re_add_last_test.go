package list_of_speakers

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

func TestReAddLastCorrect(t *testing.T) {
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
			"end_time":            1000300,
		},
	})

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeUpdate, "speaker/1")
	model := tc.GetModel("speaker/1")
	if model["begin_time"] != nil {
		t.Error("expected begin_time to be cleared after re-add")
	}
	if model["end_time"] != nil {
		t.Error("expected end_time to be cleared after re-add")
	}
}

func TestReAddLastPicksLastFinished(t *testing.T) {
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
			"end_time":            1000100,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"begin_time":          1000200,
			"end_time":            1000300,
		},
	})

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	// Speaker 2 has a later end_time, so it should be re-added.
	tc.AssertHasEvent(resp, event.TypeUpdate, "speaker/2")
	model2 := tc.GetModel("speaker/2")
	if model2["begin_time"] != nil {
		t.Error("expected speaker/2 begin_time to be cleared")
	}
	if model2["end_time"] != nil {
		t.Error("expected speaker/2 end_time to be cleared")
	}
	// Speaker 1 should remain unchanged.
	model1 := tc.GetModel("speaker/1")
	if model1["end_time"] != 1000100 {
		t.Errorf("expected speaker/1 end_time to remain 1000100, got %v", model1["end_time"])
	}
}

func TestReAddLastNoFinishedSpeakers(t *testing.T) {
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

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 1,
	})
	tc.AssertError(resp, err)
}

func TestReAddLastNoSpeakers(t *testing.T) {
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

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 1,
	})
	tc.AssertError(resp, err)
}

func TestReAddLastWrongID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 9999,
	})
	tc.AssertError(resp, err)
}

func TestReAddLastWithCurrentSpeaker(t *testing.T) {
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
			"end_time":            1000100,
		},
		"speaker/2": {
			"meeting_id":          1,
			"list_of_speakers_id": 1,
			"meeting_user_id":     1,
			"begin_time":          1000200,
		},
	})

	resp, err := tc.Request("list_of_speakers.re_add_last", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	// Speaker 1 is the only finished speaker, so it should be re-added.
	tc.AssertHasEvent(resp, event.TypeUpdate, "speaker/1")
	model1 := tc.GetModel("speaker/1")
	if model1["begin_time"] != nil {
		t.Error("expected speaker/1 begin_time to be cleared")
	}
	if model1["end_time"] != nil {
		t.Error("expected speaker/1 end_time to be cleared")
	}
}
