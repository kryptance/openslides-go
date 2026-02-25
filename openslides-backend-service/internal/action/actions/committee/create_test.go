package committee

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCommitteeCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":            "test_committee",
		"organization_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateWithDescription(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
		"committee/1": {
			"organization_id": 1,
			"name":            "c1",
		},
		"organization_tag/12": {
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":                 "test_committee2",
		"organization_id":     1,
		"description":         "<p>Test Committee</p>",
		"organization_tag_ids": []any{12},
		"forward_to_committee_ids": []any{1},
		"external_id":         "external",
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateOnlyRequired(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.create", map[string]any{
		"name":            "test_committee1",
		"organization_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateWithManagerIds(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/13": {
			"username": "test",
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":            "test_committee1",
		"organization_id": 1,
		"manager_ids":     []any{13},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCommitteeCreateMissingName(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.create", map[string]any{
		"organization_id": 1,
	})
	tc.AssertError(resp, err)
}

func TestCommitteeCreateMissingOrgId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.create", map[string]any{
		"name": "test_committee",
	})
	tc.AssertError(resp, err)
}

func TestCommitteeCreateWithExternalId(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":            "test_committee",
		"organization_id": 1,
		"external_id":     "ext_123",
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateWithForwardToCommittees(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
		"committee/1": {
			"organization_id": 1,
			"name":            "c1",
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":                     "test_committee2",
		"organization_id":         1,
		"forward_to_committee_ids": []any{1},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeCreateWithOrganizationTags(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "test_organization1",
		},
		"organization_tag/12": {
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("committee.create", map[string]any{
		"name":                 "test_committee",
		"organization_id":     1,
		"organization_tag_ids": []any{12},
	})
	tc.AssertSuccess(resp, err)
}
