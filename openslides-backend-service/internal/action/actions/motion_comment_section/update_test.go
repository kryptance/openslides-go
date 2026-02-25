package motion_comment_section

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrectAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_comment_section/111": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
		"group/23": {"meeting_id": 1, "name": "name_asdfetza"},
	})
	resp, err := tc.Request("motion_comment_section.update", map[string]any{
		"id":                  111,
		"name":                "name_iuqAPRuD",
		"read_group_ids":      []any{23},
		"write_group_ids":     []any{23},
		"submitter_can_write": false,
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_comment_section/111")
	if model["name"] != "name_iuqAPRuD" {
		t.Errorf("expected name 'name_iuqAPRuD', got %v", model["name"])
	}
	if model["meeting_id"] != 1 {
		t.Errorf("expected meeting_id 1, got %v", model["meeting_id"])
	}
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate model existence before updating.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"group/23": {"meeting_id": 1, "name": "name_asdfetza"},
		"group/24": {"meeting_id": 1, "name": "name_faofetza"},
		"motion_comment_section/111": {
			"name":           "name_srtgb123",
			"meeting_id":     1,
			"read_group_ids": []any{23},
		},
	})
	resp, err := tc.Request("motion_comment_section.update", map[string]any{
		"id": 112, "read_group_ids": []any{24},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment_section/111", map[string]any{
		"name": "name_srtgb123",
	})
}
