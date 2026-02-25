package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestResetCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "finished",
			"content_object_id": "motion/1",
			"votesvalid":        "10.000000",
			"votesinvalid":      "2.000000",
			"votescast":         "12.000000",
		},
	})

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "created",
	})
}

func TestResetClearsVoteFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":              1,
			"title":                   "Test Poll",
			"type":                    "named",
			"pollmethod":              "YNA",
			"state":                   "finished",
			"content_object_id":       "motion/1",
			"votesvalid":              "10.000000",
			"votesinvalid":            "2.000000",
			"votescast":               "12.000000",
			"entitled_users_at_stop":  "some_data",
		},
	})

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("poll/1")
	if model["votesvalid"] != nil {
		t.Errorf("expected votesvalid to be nil, got %v", model["votesvalid"])
	}
	if model["votesinvalid"] != nil {
		t.Errorf("expected votesinvalid to be nil, got %v", model["votesinvalid"])
	}
	if model["votescast"] != nil {
		t.Errorf("expected votescast to be nil, got %v", model["votescast"])
	}
	if model["entitled_users_at_stop"] != nil {
		t.Errorf("expected entitled_users_at_stop to be nil, got %v", model["entitled_users_at_stop"])
	}
}

func TestResetMotionPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
			"poll_ids":   []any{1},
		},
		"poll/1": {
			"meeting_id":        1,
			"title":             "Motion Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "finished",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "created",
	})
}

func TestResetAssignmentPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/1": {
			"meeting_id": 1,
			"title":      "Test Assignment",
			"poll_ids":   []any{1},
		},
		"poll/1": {
			"meeting_id":        1,
			"title":             "Assignment Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "started",
			"content_object_id": "assignment/1",
		},
	})

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "created",
	})
}

func TestResetFromStartedState(t *testing.T) {
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

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "created",
	})
}

func TestResetFromPublishedState(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "YNA",
			"state":             "published",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.reset", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"state": "created",
	})
}
