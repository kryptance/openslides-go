package motion_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/111": {"name": "name_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("motion_category.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion_category/111")
}

func TestDeleteWrongID(t *testing.T) {
	// The Go backend does not validate model existence before deleting.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/112": {"name": "name_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("motion_category.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_category/112")
}
