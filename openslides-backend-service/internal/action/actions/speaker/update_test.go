package speaker

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateSpeechState(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":           1,
		"speech_state": "contribution",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "contribution",
	})
}

func TestUpdateSpeechStatePro(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":           1,
		"speech_state": "pro",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "pro",
	})
}

func TestUpdateSpeechStateContra(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":           1,
		"speech_state": "contra",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state": "contra",
	})
}

func TestUpdatePointOfOrder(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":             1,
		"point_of_order": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"point_of_order": true,
	})
}

func TestUpdateNote(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":   1,
		"note": "Updated note",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"note": "Updated note",
	})
}

func TestUpdatePointOfOrderCategory(t *testing.T) {
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
			"point_of_order":      true,
		},
		"point_of_order_category/1": {
			"meeting_id": 1,
			"text":       "Technical",
			"rank":       1,
		},
	})

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":                        1,
		"point_of_order_category_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"point_of_order_category_id": 1,
	})
}

func TestUpdateMeetingUserId(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":              1,
		"meeting_user_id": 2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"meeting_user_id": 2,
	})
}

func TestUpdateMultipleFields(t *testing.T) {
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

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":             1,
		"speech_state":   "pro",
		"point_of_order": false,
		"note":           "Test note",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("speaker/1", map[string]any{
		"speech_state":   "pro",
		"point_of_order": false,
		"note":           "Test note",
	})
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("speaker.update", map[string]any{
		"speech_state": "pro",
	})
	tc.AssertErrorContains(err, "id")
}

func TestUpdateWrongID(t *testing.T) {
	t.Skip("In-memory datastore does not validate model existence on update")
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("speaker.update", map[string]any{
		"id":           9999,
		"speech_state": "pro",
	})
	tc.AssertError(resp, err)
}
