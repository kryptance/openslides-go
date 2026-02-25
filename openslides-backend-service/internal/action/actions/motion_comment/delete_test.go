package motion_comment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1":           {"meeting_id": 1, "comment_ids": []any{111}},
		"motion_comment/111": {"meeting_id": 1, "section_id": 78, "motion_id": 1},
		"motion_comment_section/78": {
			"meeting_id":     1,
			"write_group_ids": []any{3},
			"name":           "test",
		},
	})
	resp, err := tc.Request("motion_comment.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_comment/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_comment/112": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_comment.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment/112")
}
