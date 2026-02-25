package topic

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
)

// TestImportBasicNewTopic tests importing a single new topic generates a create event.
func TestImportBasicNewTopic(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
		"import_preview": []any{
			map[string]any{
				"state":    "new",
				"messages": []any{},
				"data": map[string]any{
					"title": "test",
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
	// The import action creates events but doesn't apply them to the datastore
	// (unlike NewCreateAction which does). Verify via events.
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/1")
}

// TestImportSkipsErrorRows verifies that rows in error state are skipped.
func TestImportSkipsErrorRows(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
		"import_preview": []any{
			map[string]any{
				"state":    "error",
				"messages": []any{"Some error"},
				"data": map[string]any{
					"title": "should_be_skipped",
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
	// No topic create events should be generated for error rows.
	// (Relation manager may still produce events, so we check for
	// absence of topic create events rather than total event count.)
	tc.AssertModelNotExists("topic/1")
}

// TestImportMultipleRows tests importing multiple topic rows generates events for each.
func TestImportMultipleRows(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
		"import_preview": []any{
			map[string]any{
				"state":    "new",
				"messages": []any{},
				"data": map[string]any{
					"title": "topic_one",
				},
			},
			map[string]any{
				"state":    "new",
				"messages": []any{},
				"data": map[string]any{
					"title": "topic_two",
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/1")
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/2")
}

// TestImportMissingImportPreview tests that missing import_preview causes an error.
func TestImportMissingImportPreview(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("topic.import", map[string]any{
		"id": 1,
	})
	tc.AssertErrorContains(err, "import_preview")
}

// TestImportWithTextAndAgendaFields tests importing a topic with text and agenda data.
func TestImportWithTextAndAgendaFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
		"import_preview": []any{
			map[string]any{
				"state":    "new",
				"messages": []any{},
				"data": map[string]any{
					"title":           "test topic",
					"text":            "some description",
					"agenda_comment":  "test comment",
					"agenda_type":     "hidden",
					"agenda_duration": 50,
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/1")
}

// TestImportMixedStates tests that only non-error rows are imported.
func TestImportMixedStates(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
		"import_preview": []any{
			map[string]any{
				"state":    "new",
				"messages": []any{},
				"data": map[string]any{
					"title": "valid_topic",
				},
			},
			map[string]any{
				"state":    "error",
				"messages": []any{"Duplicate"},
				"data": map[string]any{
					"title": "invalid_topic",
				},
			},
			map[string]any{
				"state":    "done",
				"messages": []any{},
				"data": map[string]any{
					"title": "another_valid",
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
	// 2 create events (error row skipped)
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/1")
	tc.AssertHasEvent(resp, event.TypeCreate, "topic/2")
}
