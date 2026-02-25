package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSetStateAnalogStartToFinished(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Analog Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	// Start the poll.
	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})

	// Stop the poll.
	resp, err = tc.Request("poll.stop", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "finished",
	})

	// Publish the poll.
	resp, err = tc.Request("poll.publish", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "published",
	})
}

func TestSetStateFullLifecycle(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Lifecycle Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	// Start.
	resp, err := tc.Request("poll.start", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{"state": "started"})

	// Stop.
	resp, err = tc.Request("poll.stop", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{"state": "finished"})

	// Publish.
	resp, err = tc.Request("poll.publish", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{"state": "published"})

	// Reset.
	resp, err = tc.Request("poll.reset", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{"state": "created"})
}
