package point_of_order_category

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("point_of_order_category.create", map[string]any{
		"text":       "blablabla",
		"rank":       11,
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("point_of_order_category/1", map[string]any{
		"text":       "blablabla",
		"rank":       11,
		"meeting_id": 1,
	})
}
