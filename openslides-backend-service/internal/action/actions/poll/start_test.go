package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestStartCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartAnalogPoll(t *testing.T) {
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

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartNamedPollYNA(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Named YNA Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartNamedPollY(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Named Y Poll",
			"type":              "named",
			"pollmethod":        "Y",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartNamedPollN(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Named N Poll",
			"type":              "named",
			"pollmethod":        "N",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartPseudoanonymousPollYNA(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Pseudoanonymous YNA Poll",
			"type":              "pseudoanonymous",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartPseudoanonymousPollY(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Pseudoanonymous Y Poll",
			"type":              "pseudoanonymous",
			"pollmethod":        "Y",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}

func TestStartPseudoanonymousPollN(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Pseudoanonymous N Poll",
			"type":              "pseudoanonymous",
			"pollmethod":        "N",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.start", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "started",
	})
}
