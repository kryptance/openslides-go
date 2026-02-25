package motion_block

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_block/111": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("motion_block.update", map[string]any{
		"id": 111, "title": "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_block/111", map[string]any{"title": "title_Xcdfgee"})
}

func TestUpdateWrongID(t *testing.T) {
	// The Go backend does not validate model existence before updating.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_block/111": {"title": "title_srtgb123", "meeting_id": 1},
	})
	resp, err := tc.Request("motion_block.update", map[string]any{
		"id": 112, "title": "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_block/111", map[string]any{"title": "title_srtgb123"})
}
