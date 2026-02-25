package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestStopCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "started",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.stop", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "finished",
	})
}

func TestStopNamedPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Named Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "started",
			"content_object_id": "motion/1",
			"entitled_group_ids": []any{1},
		},
	})

	resp, err := tc.Request("poll.stop", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "finished",
	})
}

func TestStopPseudoanonymousPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Pseudoanonymous Poll",
			"type":              "pseudoanonymous",
			"pollmethod":        "YNA",
			"state":             "started",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.stop", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "finished",
	})
}

func TestStopAnalogPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Analog Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "started",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.stop", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "finished",
	})
}
