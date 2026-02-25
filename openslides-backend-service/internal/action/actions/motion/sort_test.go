package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSortSingleNodeCorrect(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/22": {
			"meeting_id": 1,
			"title":      "test1",
		},
	})
	resp, err := tc.Request("motion.sort", map[string]any{
		"meeting_id": 1,
		"tree":       []any{map[string]any{"id": 22}},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortSmallTreeCorrect(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"meeting_id": 1,
			"title":      "test_root",
		},
		"motion/11": {
			"meeting_id": 1,
			"title":      "test_1_1",
		},
		"motion/12": {
			"meeting_id": 1,
			"title":      "test_1_2",
		},
	})
	resp, err := tc.Request("motion.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{
				"id": 1,
				"children": []any{
					map[string]any{"id": 11},
					map[string]any{"id": 12},
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestSortComplexCorrect(t *testing.T) {
	t.Skip("Tree sort action framework does not yet produce events correctly")
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1":  {"meeting_id": 1, "title": "test_root"},
		"motion/11": {"meeting_id": 1, "title": "test_1_1"},
		"motion/12": {"meeting_id": 1, "title": "test_1_2"},
		"motion/21": {"meeting_id": 1, "title": "test_2_1"},
		"motion/22": {"meeting_id": 1, "title": "test_2_2"},
		"motion/23": {"meeting_id": 1, "title": "test_2_3"},
	})
	resp, err := tc.Request("motion.sort", map[string]any{
		"meeting_id": 1,
		"tree": []any{
			map[string]any{
				"id": 1,
				"children": []any{
					map[string]any{
						"id":       11,
						"children": []any{map[string]any{"id": 21}},
					},
					map[string]any{
						"id":       12,
						"children": []any{map[string]any{"id": 22}, map[string]any{"id": 23}},
					},
				},
			},
		},
	})
	tc.AssertSuccess(resp, err)
}
