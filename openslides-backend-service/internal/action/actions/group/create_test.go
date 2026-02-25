package group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/22": {
			"name":                         "name_vJxebUwo",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{22}},
	})

	resp, err := tc.Request("group.create", map[string]any{
		"name":       "test_Xcdfgee",
		"meeting_id": 22,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("group/1")
	if model["name"] != "test_Xcdfgee" {
		t.Errorf("expected name test_Xcdfgee, got %v", model["name"])
	}
	if model["meeting_id"] != 22 {
		t.Errorf("expected meeting_id 22, got %v", model["meeting_id"])
	}
}

func TestCreateWithPermissions(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/22": {
			"name":                         "name_vJxebUwo",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{22}},
	})

	resp, err := tc.Request("group.create", map[string]any{
		"name":        "test_Xcdfgee",
		"meeting_id":  22,
		"permissions": []any{"agenda_item.can_see"},
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("group/1")
	if model["name"] != "test_Xcdfgee" {
		t.Errorf("expected name test_Xcdfgee, got %v", model["name"])
	}
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/22": {
			"name":                         "name_vJxebUwo",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{22}},
	})

	_, err := tc.Request("group.create", map[string]any{
		"meeting_id": 22,
	})
	tc.AssertErrorContains(err, "name")
}

func TestCreateExternalID(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"committee/1": {"meeting_ids": []any{22}},
		"meeting/22": {
			"name":                         "name_vJxebUwo",
			"admin_group_id":               3,
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"group/3": {
			"name":                       "test",
			"admin_group_for_meeting_id": 22,
			"meeting_id":                 22,
		},
	})

	resp, err := tc.Request("group.create", map[string]any{
		"name":        "test_name",
		"external_id": "test",
		"meeting_id":  22,
	})
	tc.AssertSuccess(resp, err)
	// SetModels with group/3 sets nextID["group"]=3, so the next create gets group/4.
	tc.AssertModelExists("group/4", map[string]any{
		"external_id": "test",
		"name":        "test_name",
		"meeting_id":  22,
	})
}
