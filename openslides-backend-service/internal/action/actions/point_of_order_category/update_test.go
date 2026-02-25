package point_of_order_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"point_of_order_category_ids": []any{45}},
		"point_of_order_category/45": {
			"text":       "blablabla",
			"rank":       11,
			"meeting_id": 1,
		},
	})

	resp, err := tc.Request("point_of_order_category.update", map[string]any{
		"id":   45,
		"text": "foo",
		"rank": 12,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("point_of_order_category/45", map[string]any{
		"text": "foo",
		"rank": 12,
	})
}
