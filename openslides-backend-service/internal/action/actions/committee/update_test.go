package committee

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupUpdateCommittee(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
		"committee/1": {
			"name":            "committee_testname",
			"description":     "<p>Test description</p>",
			"organization_id": 1,
		},
		"committee/2": {
			"name":            "forwarded_committee",
			"organization_id": 1,
		},
		"user/20": {
			"username": "test_user20",
		},
		"user/21": {
			"username": "test_user21",
		},
	})
}

func TestCommitteeUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":   1,
		"name": "committee_testname_updated",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("committee/1", map[string]any{
		"name": "committee_testname_updated",
	})
}

func TestCommitteeUpdateDescription(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":          1,
		"description": "<p>New Test description</p>",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("committee/1", map[string]any{
		"description": "<p>New Test description</p>",
	})
}

func TestCommitteeUpdateExternalId(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":          1,
		"external_id": "ext_123",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("committee/1", map[string]any{
		"external_id": "ext_123",
	})
}

func TestCommitteeUpdateManagerIds(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":          1,
		"manager_ids": []any{20, 21},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeUpdateForwardToCommittees(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":                       1,
		"forward_to_committee_ids": []any{2},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeUpdateOrganizationTags(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	tc.SetModels(map[string]map[string]any{
		"organization_tag/12": {
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("committee.update", map[string]any{
		"id":                   1,
		"organization_tag_ids": []any{12},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeUpdateMissingId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.update", map[string]any{
		"name": "test",
	})
	tc.AssertError(resp, err)
}

func TestCommitteeUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	setupUpdateCommittee(tc)
	resp, err := tc.Request("committee.update", map[string]any{
		"id":          1,
		"name":        "new_name",
		"description": "new description",
		"external_id": "ext_456",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("committee/1", map[string]any{
		"name":        "new_name",
		"description": "new description",
		"external_id": "ext_456",
	})
}
