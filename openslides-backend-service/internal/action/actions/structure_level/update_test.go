package structure_level

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"structure_level_ids": []any{1},
		},
		"structure_level/1": {"meeting_id": 1, "name": "test"},
	})

	resp, err := tc.Request("structure_level.update", map[string]any{
		"id":           1,
		"name":         "test2",
		"color":        "#abf257",
		"default_time": 600,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("structure_level/1", map[string]any{
		"name":         "test2",
		"meeting_id":   1,
		"color":        "#abf257",
		"default_time": 600,
	})
}
