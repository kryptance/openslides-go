package motion

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"title":      "title_srtgb123",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion/111")
}

func TestDeleteAmendment(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"title":         "title_srtgb123",
			"meeting_id":    1,
			"amendment_ids": []any{222},
		},
		"motion/222": {
			"title":          "amendment to 111",
			"meeting_id":     1,
			"lead_motion_id": 111,
		},
	})
	resp, err := tc.Request("motion.delete", map[string]any{"id": 222})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion/111")
	tc.AssertModelDeleted("motion/222")
}

func TestDeleteMotionAndAmendment(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/111": {
			"title":         "title_srtgb123",
			"meeting_id":    1,
			"amendment_ids": []any{222},
		},
		"motion/222": {
			"title":          "amendment to 111",
			"meeting_id":     1,
			"lead_motion_id": 111,
		},
	})
	resp, err := tc.RequestMulti("motion.delete", []map[string]any{
		{"id": 111},
		{"id": 222},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("motion/111")
	tc.AssertModelDeleted("motion/222")
}
