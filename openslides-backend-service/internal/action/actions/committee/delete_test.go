package committee

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupDeleteCommittee(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"committee_ids": []any{1},
		},
		"user/20": {
			"committee_ids": []any{1},
		},
		"user/21": {
			"committee_ids":            []any{1},
			"committee_management_ids": []any{1},
		},
		"committee/1": {
			"organization_id": 1,
			"user_ids":        []any{20, 21},
			"manager_ids":     []any{21},
		},
	})
}

func TestCommitteeDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	setupDeleteCommittee(tc)
	resp, err := tc.Request("committee.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("committee/1")
}

func TestCommitteeDeleteWithForwardings(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"committee_ids": []any{1, 2, 3},
		},
		"committee/1": {
			"organization_id":                       1,
			"organization_tag_ids":                   []any{12},
			"forward_to_committee_ids":               []any{2},
			"receive_forwardings_from_committee_ids": []any{3},
		},
		"committee/2": {
			"organization_id":                       1,
			"receive_forwardings_from_committee_ids": []any{1},
		},
		"committee/3": {
			"organization_id":         1,
			"forward_to_committee_ids": []any{1},
		},
		"organization_tag/12": {
			"tagged_ids":      []any{"committee/1"},
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("committee.delete", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("committee/1")
}

func TestCommitteeDeleteMissingId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.delete", map[string]any{})
	tc.AssertError(resp, err)
}
