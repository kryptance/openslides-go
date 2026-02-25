package motion_comment_section

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_comment_section/111": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_comment_section.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_comment_section/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_comment_section/112": {
			"name":       "name_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_comment_section.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment_section/112")
}
