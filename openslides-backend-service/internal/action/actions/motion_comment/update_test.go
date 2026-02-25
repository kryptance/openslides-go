package motion_comment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {"meeting_id": 1, "comment_ids": []any{111}},
		"motion_comment/111": {
			"comment":    "comment_srtgb123",
			"meeting_id": 1,
			"motion_id":  111,
			"section_id": 78,
		},
		"motion_comment_section/78": {
			"meeting_id":     1,
			"write_group_ids": []any{3},
			"name":           "test",
		},
	})
	resp, err := tc.Request("motion_comment.update", map[string]any{
		"id": 111, "comment": "comment_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("motion_comment/111")
	if model["comment"] != "comment_Xcdfgee" {
		t.Errorf("expected comment 'comment_Xcdfgee', got %v", model["comment"])
	}
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate model existence before updating.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_comment/111": {
			"comment":    "comment_srtgb123",
			"meeting_id": 1,
			"section_id": 78,
		},
		"motion_comment_section/78": {"meeting_id": 1, "write_group_ids": []any{3}},
	})
	resp, err := tc.Request("motion_comment.update", map[string]any{
		"id": 112, "comment": "comment_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment/111", map[string]any{"comment": "comment_srtgb123"})
}
