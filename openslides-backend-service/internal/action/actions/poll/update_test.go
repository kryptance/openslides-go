package poll

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":    1,
		"title": "Updated Title",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title": "Updated Title",
	})
}

func TestUpdateDescription(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":          1,
		"description": "New description",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"description": "New description",
	})
}

func TestUpdateGlobalYes(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":         1,
		"global_yes": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"global_yes": true,
	})
}

func TestUpdateGlobalNo(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":        1,
		"global_no": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"global_no": true,
	})
}

func TestUpdateGlobalAbstain(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":             1,
		"global_abstain": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"global_abstain": true,
	})
}

func TestUpdateMinVotesAmount(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "Y",
			"state":             "created",
			"content_object_id": "assignment/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":               1,
		"min_votes_amount": 2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"min_votes_amount": 2,
	})
}

func TestUpdateMaxVotesAmount(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "Y",
			"state":             "created",
			"content_object_id": "assignment/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":               1,
		"max_votes_amount": 10,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"max_votes_amount": 10,
	})
}

func TestUpdateMaxVotesPerOption(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "named",
			"pollmethod":        "Y",
			"state":             "created",
			"content_object_id": "assignment/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":                   1,
		"max_votes_per_option": 3,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"max_votes_per_option": 3,
	})
}

func TestUpdateOnehundredPercentBase(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":                      1,
		"onehundred_percent_base": "YN",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"onehundred_percent_base": "YN",
	})
}

func TestUpdateEntitledGroupIds(t *testing.T) {
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

	resp, err := tc.Request("poll.update", map[string]any{
		"id":                 1,
		"entitled_group_ids": []any{1, 2},
	})
	tc.AssertSuccess(resp, err)
}

func TestUpdateBackend(t *testing.T) {
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
			"backend":           "long",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":      1,
		"backend": "fast",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"backend": "fast",
	})
}

func TestUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"poll/1": {
			"meeting_id":        1,
			"title":             "Test Poll",
			"type":              "analog",
			"pollmethod":        "YNA",
			"state":             "created",
			"content_object_id": "motion/1",
		},
	})

	resp, err := tc.Request("poll.update", map[string]any{
		"id":          1,
		"title":       "Updated",
		"description": "Updated description",
		"global_yes":  true,
		"global_no":   false,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("poll/1", map[string]any{
		"title":       "Updated",
		"description": "Updated description",
		"global_yes":  true,
		"global_no":   false,
	})
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	_, err := tc.Request("poll.update", map[string]any{
		"title": "Updated",
	})
	tc.AssertErrorContains(err, "id")
}
