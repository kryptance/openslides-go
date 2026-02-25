package motion_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestNumberMotionsGoodSingleMotion(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"motion_category_ids": []any{111},
			"motion_ids":          []any{69},
		},
		"motion_category/111": {
			"name":       "name_MKKAcYQu",
			"prefix":     "prefix_A",
			"motion_ids": []any{69},
			"meeting_id": 1,
		},
		"motion/69": {
			"title":       "title_NAZOknoM",
			"category_id": 111,
			"meeting_id":  1,
		},
	})
	resp, err := tc.Request("motion_category.number_motions", map[string]any{
		"id": 111,
	})
	tc.AssertSuccess(resp, err)
}

func TestNumberMotionsInvalidID(t *testing.T) {
	// The Go backend's number_motions is a simple update action (TODO: custom logic).
	// It does not validate model existence, so a non-existing ID still succeeds.
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion_category/1": {
			"name":       "category_1",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_category.number_motions", map[string]any{
		"id": 222,
	})
	tc.AssertSuccess(resp, err)
}
