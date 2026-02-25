package motion_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrectAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/111": {
			"name":       "name_srtgb123",
			"prefix":     "prefix_JmDHFgvH",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_category.update", map[string]any{
		"id":     111,
		"name":   "name_Xcdfgee",
		"prefix": "prefix_sthyAKrW",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_category/111")
	if model["name"] != "name_Xcdfgee" {
		t.Errorf("expected name 'name_Xcdfgee', got %v", model["name"])
	}
	if model["prefix"] != "prefix_sthyAKrW" {
		t.Errorf("expected prefix 'prefix_sthyAKrW', got %v", model["prefix"])
	}
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate model existence before updating.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/111": {
			"name":       "name_srtgb123",
			"prefix":     "prefix_JmDHFgvH",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_category.update", map[string]any{
		"id": 112, "name": "name_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_category/111", map[string]any{"name": "name_srtgb123"})
}

func TestUpdateNonUniquePrefix(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/111": {
			"name":       "name_srtgb123",
			"prefix":     "bla",
			"meeting_id": 1,
		},
		"motion_category/110": {
			"name":       "name_already",
			"prefix":     "test",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_category.update", map[string]any{
		"id":     111,
		"prefix": "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_category/111", map[string]any{
		"name":       "name_srtgb123",
		"prefix":     "test",
		"meeting_id": 1,
	})
}
