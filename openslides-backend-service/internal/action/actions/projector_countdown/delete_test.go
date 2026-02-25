package projector_countdown

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDelete(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"projector_countdown_ids": []any{1},
		},
		"projector_countdown/1": {"meeting_id": 1, "title": "test1"},
	})

	resp, err := tc.Request("projector_countdown.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("projector_countdown/1")
}
