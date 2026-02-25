package motion_block

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_block/111": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_block.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_block/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_block/112": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_block.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_block/112")
}

func TestDeleteCorrectCascading(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_block/111": {
			"list_of_speakers_id": 222,
			"agenda_item_id":      333,
			"meeting_id":          1,
		},
		"list_of_speakers/222": {
			"closed":            false,
			"content_object_id": "motion_block/111",
			"meeting_id":        1,
		},
		"agenda_item/333": {
			"comment":           "test_comment_ewoirzewoirioewr",
			"content_object_id": "motion_block/111",
			"meeting_id":        1,
		},
	})
	resp, err := tc.Request("motion_block.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_block/111")
}
