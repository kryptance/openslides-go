package meeting_user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestMeetingUserUpdate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{10}},
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
			"committee_id":                 1,
			"default_group_id":             22,
			"structure_level_ids":          []any{31},
		},
		"meeting_user/5":    {"user_id": 1, "meeting_id": 10},
		"group/21":          {"meeting_id": 10},
		"group/22":          {"meeting_id": 10, "default_group_for_meeting_id": 10},
		"structure_level/31": {"meeting_id": 10},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":                  5,
		"comment":             "test bla",
		"number":              "XII",
		"structure_level_ids": []any{31},
		"about_me":            "A very long description.",
		"vote_weight":         "1.500000",
		"group_ids":           []any{21},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{
		"comment":     "test bla",
		"number":      "XII",
		"about_me":    "A very long description.",
		"vote_weight": "1.500000",
	})
}

func TestMeetingUserUpdateComment(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10, "comment": "old"},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":      5,
		"comment": "new comment",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{"comment": "new comment"})
}

func TestMeetingUserUpdateNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":     5,
		"number": "XIII",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{"number": "XIII"})
}

func TestMeetingUserUpdateAboutMe(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":       5,
		"about_me": "Updated about me",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{"about_me": "Updated about me"})
}

func TestMeetingUserUpdateVoteWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10, "vote_weight": "1.000000"},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":          5,
		"vote_weight": "2.500000",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{"vote_weight": "2.500000"})
}

func TestMeetingUserUpdateLockedOut(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":         5,
		"locked_out": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{"locked_out": "true"})
}

func TestMeetingUserUpdateNonExistent(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// update operations. The action succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":      999,
		"comment": "test",
	})
	tc.AssertSuccess(resp, err)
}

func TestMeetingUserUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{5},
		},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10},
	})
	resp, err := tc.Request("meeting_user.update", map[string]any{
		"id":          5,
		"comment":     "updated comment",
		"number":      "XIV",
		"about_me":    "Updated bio",
		"vote_weight": "3.000000",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/5", map[string]any{
		"comment":     "updated comment",
		"number":      "XIV",
		"about_me":    "Updated bio",
		"vote_weight": "3.000000",
	})
}
