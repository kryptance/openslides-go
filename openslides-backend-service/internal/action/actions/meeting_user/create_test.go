package meeting_user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestMeetingUserCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{10}},
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
			"structure_level_ids":          []any{31},
			"group_ids":                    []any{21},
		},
		"group/21":           {"meeting_id": 10},
		"structure_level/31": {"meeting_id": 10},
		"user/1":             {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":             1,
		"meeting_id":          10,
		"comment":             "test blablaba",
		"number":              "XII",
		"structure_level_ids": []any{31},
		"about_me":            "A very long description.",
		"vote_weight":         "1.500000",
		"group_ids":           []any{21},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"user_id":     "1",
		"meeting_id":  "10",
		"comment":     "test blablaba",
		"number":      "XII",
		"about_me":    "A very long description.",
		"vote_weight": "1.500000",
	})
}

func TestMeetingUserCreateMinimal(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":    1,
		"meeting_id": 10,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"user_id":    "1",
		"meeting_id": "10",
	})
}

func TestMeetingUserCreateWithComment(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":    1,
		"meeting_id": 10,
		"comment":    "test comment",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"comment": "test comment",
	})
}

func TestMeetingUserCreateWithNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":    1,
		"meeting_id": 10,
		"number":     "42",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"number": "42",
	})
}

func TestMeetingUserCreateWithAboutMe(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":    1,
		"meeting_id": 10,
		"about_me":   "<p>About me text</p>",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"about_me": "<p>About me text</p>",
	})
}

func TestMeetingUserCreateWithVoteWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":     1,
		"meeting_id":  10,
		"vote_weight": "2.500000",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"vote_weight": "2.500000",
	})
}

func TestMeetingUserCreateDifferentUser(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{10}},
		"user/1":      {"username": "admin"},
		"user/5":      {"username": "testuser"},
	})
	resp, err := tc.Request("meeting_user.create", map[string]any{
		"user_id":    5,
		"meeting_id": 10,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/1", map[string]any{
		"user_id":    "5",
		"meeting_id": "10",
	})
}
