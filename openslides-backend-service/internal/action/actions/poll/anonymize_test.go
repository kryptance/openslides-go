package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestAnonymizeCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":  1,
			"title":       "Test Poll",
			"type":        "named",
			"pollmethod":  "YNA",
			"state":       "finished",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.anonymize", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"is_pseudoanonymized": true,
	})
}

func TestAnonymizeSetsPseudoanonymized(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":  1,
			"title":       "Test Poll",
			"type":        "named",
			"pollmethod":  "YNA",
			"state":       "published",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.anonymize", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("poll/1")
	if model["is_pseudoanonymized"] != true {
		t.Errorf("expected is_pseudoanonymized to be true, got %v", model["is_pseudoanonymized"])
	}
}

func TestAnonymizeMotionPoll(t *testing.T) {
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

	resp, err := tc.Request("poll.anonymize", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"is_pseudoanonymized": true,
	})
}

func TestAnonymizeAssignmentPoll(t *testing.T) {
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
			"state":             "finished",
			"content_object_id": "assignment/1",
		},
	})

	resp, err := tc.Request("poll.anonymize", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"is_pseudoanonymized": true,
	})
}
