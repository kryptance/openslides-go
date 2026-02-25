package motion_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortMotionsCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/222": {"meeting_id": 1},
		"motion/31":           {"category_id": 222, "meeting_id": 1},
		"motion/32":           {"category_id": 222, "meeting_id": 1},
	})
	resp, err := tc.Request("motion_category.sort_motions_in_category", map[string]any{
		"id": 222, "motion_ids": []any{32, 31},
	})
	tc.AssertSuccess(resp, err)
}
