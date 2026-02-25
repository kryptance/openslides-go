package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"name":         "test_name",
		"committee_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithDescription(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"name":         "test_name",
		"committee_id": 1,
		"description":  "RRfnzxHA",
		"location":     "LSFHPTgE",
		"start_time":   1608120653,
		"end_time":     1608121653,
		"external_id":  "external",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithWelcomeFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"name":          "test_name",
		"committee_id":  1,
		"welcome_title": "Welcome",
		"welcome_text":  "Hello everyone",
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateWithOrganizationTags(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
		"organization_tag/3": {
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"name":                 "test_name",
		"committee_id":         1,
		"organization_tag_ids": []any{3},
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateMissingName(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
		"committee/1": {
			"name":            "test_committee",
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"committee_id": 1,
	})
	tc.AssertError(resp, err)
}

func TestCreateMissingCommittee(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"limit_of_meetings": 0,
			"active_meeting_ids": []any{},
		},
	})
	resp, err := tc.Request("meeting.create", map[string]any{
		"name": "test_name",
	})
	tc.AssertError(resp, err)
}
