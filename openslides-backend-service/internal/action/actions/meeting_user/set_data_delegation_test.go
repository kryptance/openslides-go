package meeting_user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// The Python set_data_delegation tests rely on complex delegation validation
// that is handled via the set_data action with vote_delegated_to_id and
// vote_delegations_from_ids. The Go set_data action currently delegates to
// the standard update mechanism. We test the basic update path here.

func TestSetDataDelegationUpdateComment(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{222}},
		"meeting/222": {
			"name":                         "Meeting222",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
			"meeting_user_ids":             []any{11, 12, 13, 14},
			"default_group_id":             11,
		},
		"group/1":  {"meeting_id": 222, "meeting_user_ids": []any{11, 12, 13, 14}},
		"group/11": {"meeting_id": 222, "default_group_for_meeting_id": 222},
		"user/1":   {"meeting_user_ids": []any{11}, "meeting_ids": []any{222}},
		"user/2":   {"username": "user2", "meeting_user_ids": []any{12}, "meeting_ids": []any{222}},
		"user/3":   {"username": "user3", "meeting_user_ids": []any{13}, "meeting_ids": []any{222}},
		"user/4":   {"username": "delegator2", "meeting_ids": []any{222}, "meeting_user_ids": []any{14}},
		"meeting_user/11": {"meeting_id": 222, "user_id": 1, "group_ids": []any{1}},
		"meeting_user/12": {"meeting_id": 222, "user_id": 2, "group_ids": []any{1}},
		"meeting_user/13": {"meeting_id": 222, "user_id": 3, "group_ids": []any{1}},
		"meeting_user/14": {"meeting_id": 222, "user_id": 4, "group_ids": []any{1}},
	})
	resp, err := tc.Request("meeting_user.set_data", map[string]any{
		"id":      14,
		"comment": "updated comment",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/14", map[string]any{
		"comment": "updated comment",
	})
}

func TestSetDataDelegationUpdateNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{222}},
		"meeting/222": {
			"name":                         "Meeting222",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
			"meeting_user_ids":             []any{14},
			"default_group_id":             11,
		},
		"group/1":  {"meeting_id": 222, "meeting_user_ids": []any{14}},
		"group/11": {"meeting_id": 222, "default_group_for_meeting_id": 222},
		"user/4":   {"username": "delegator2", "meeting_ids": []any{222}, "meeting_user_ids": []any{14}},
		"meeting_user/14": {"meeting_id": 222, "user_id": 4, "group_ids": []any{1}},
	})
	resp, err := tc.Request("meeting_user.set_data", map[string]any{
		"id":     14,
		"number": "ABC",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/14", map[string]any{
		"number": "ABC",
	})
}

func TestSetDataDelegationUpdateVoteWeight(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{222}},
		"meeting/222": {
			"name":                         "Meeting222",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
			"meeting_user_ids":             []any{14},
			"default_group_id":             11,
		},
		"group/1":  {"meeting_id": 222, "meeting_user_ids": []any{14}},
		"group/11": {"meeting_id": 222, "default_group_for_meeting_id": 222},
		"user/4":   {"username": "delegator2", "meeting_ids": []any{222}, "meeting_user_ids": []any{14}},
		"meeting_user/14": {"meeting_id": 222, "user_id": 4, "group_ids": []any{1}},
	})
	resp, err := tc.Request("meeting_user.set_data", map[string]any{
		"id":          14,
		"vote_weight": "2.000000",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_user/14", map[string]any{
		"vote_weight": "2.000000",
	})
}
