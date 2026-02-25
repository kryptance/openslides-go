package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "YNA",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":      "Vote",
		"type":       "analog",
		"pollmethod": "YNA",
		"meeting_id": 1,
	})
}

func TestCreateNamedPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Named Vote",
		"type":              "named",
		"pollmethod":        "YN",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":      "Named Vote",
		"type":       "named",
		"pollmethod": "YN",
	})
}

func TestCreatePseudoanonymousPoll(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Pseudoanonymous Vote",
		"type":              "pseudoanonymous",
		"pollmethod":        "Y",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":      "Pseudoanonymous Vote",
		"type":       "pseudoanonymous",
		"pollmethod": "Y",
	})
}

func TestCreateWithDescription(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "YNA",
		"description":       "A description",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":       "Vote",
		"description": "A description",
	})
}

func TestCreateWithGlobalOptions(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "YNA",
		"global_yes":        true,
		"global_no":         true,
		"global_abstain":    true,
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"global_yes":     true,
		"global_no":      true,
		"global_abstain": true,
	})
}

func TestCreateWithEntitledGroups(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":         1,
		"title":              "Vote",
		"type":               "named",
		"pollmethod":         "YNA",
		"entitled_group_ids": []any{1, 2},
		"content_object_id":  "motion/1",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithMinMaxVotes(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/1": {
			"meeting_id": 1,
			"title":      "Test Assignment",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":           1,
		"title":                "Assignment Vote",
		"type":                 "named",
		"pollmethod":           "Y",
		"min_votes_amount":     1,
		"max_votes_amount":     5,
		"max_votes_per_option": 1,
		"content_object_id":    "assignment/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"min_votes_amount":     1,
		"max_votes_amount":     5,
		"max_votes_per_option": 1,
	})
}

func TestCreateWithBackendFast(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "named",
		"pollmethod":        "YNA",
		"backend":           "fast",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"backend": "fast",
	})
}

func TestCreateWithBackendLong(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "named",
		"pollmethod":        "YNA",
		"backend":           "long",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"backend": "long",
	})
}

func TestCreateMissingTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.create", map[string]any{
		"meeting_id": 1,
		"type":       "analog",
		"pollmethod": "YNA",
	})
	tc.AssertErrorContains(err, "title")
}

func TestCreateMissingMeetingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.create", map[string]any{
		"title":      "Vote",
		"type":       "analog",
		"pollmethod": "YNA",
	})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestCreateMissingType(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.create", map[string]any{
		"meeting_id": 1,
		"title":      "Vote",
		"pollmethod": "YNA",
	})
	tc.AssertErrorContains(err, "type")
}

func TestCreateMissingPollmethod(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.create", map[string]any{
		"meeting_id": 1,
		"title":      "Vote",
		"type":       "analog",
	})
	tc.AssertErrorContains(err, "pollmethod")
}

func TestCreatePollmethodY(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "Y",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"pollmethod": "Y",
	})
}

func TestCreatePollmethodN(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "N",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"pollmethod": "N",
	})
}

func TestCreatePollmethodYN(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Vote",
		"type":              "analog",
		"pollmethod":        "YN",
		"content_object_id": "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"pollmethod": "YN",
	})
}

func TestCreateForAssignment(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"assignment/1": {
			"meeting_id": 1,
			"title":      "Test Assignment",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":        1,
		"title":             "Assignment Vote",
		"type":              "analog",
		"pollmethod":        "YNA",
		"content_object_id": "assignment/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":             "Assignment Vote",
		"content_object_id": "assignment/1",
	})
}

func TestCreateWithOnehundredPercentBase(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "Test Motion",
		},
	})

	resp, err := tc.Request("poll.create", map[string]any{
		"meeting_id":              1,
		"title":                   "Vote",
		"type":                    "analog",
		"pollmethod":              "YNA",
		"onehundred_percent_base": "YNA",
		"content_object_id":       "motion/1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"onehundred_percent_base": "YNA",
	})
}
