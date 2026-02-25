package motion_block

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":               "Test Organization",
			"active_meeting_ids": []any{42},
			"committee_ids":      []any{1},
		},
		"committee/1": {
			"name":            "Test Committee",
			"organization_id": 1,
			"meeting_ids":     []any{42},
		},
		"meeting/42": {
			"name":                        "test",
			"agenda_item_creation":         "always",
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"user/1": {
			"username":        "admin",
			"is_active":       true,
			"organization_id": 1,
		},
	})
	resp, err := tc.Request("motion_block.create", map[string]any{
		"title": "test_Xcdfgee", "meeting_id": 42,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_block/1", map[string]any{
		"title": "test_Xcdfgee",
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_block.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateWrongField(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_block.create", map[string]any{
		"wrong_field": "text_AefohteiF8",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}
