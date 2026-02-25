package structure_level

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"structure_level/1": {"meeting_id": 1, "name": "test"},
	})

	resp, err := tc.Request("structure_level.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("structure_level/1")
}
