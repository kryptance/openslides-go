package motion_submitter

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/357": {
			"title":      "title_YIDYXmKj",
			"meeting_id": 1,
		},
		"user/78": {
			"username":    "username_loetzbfg",
			"meeting_ids": []any{1},
		},
		"meeting_user/79": {
			"meeting_id": 1,
			"user_id":    78,
		},
	})
	resp, err := tc.Request("motion_submitter.create", map[string]any{
		"motion_id":       357,
		"meeting_user_id": 79,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_submitter.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCreateWithMotionIdOnly(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/357": {
			"title":      "title_YIDYXmKj",
			"meeting_id": 1,
		},
	})
	resp, err := tc.Request("motion_submitter.create", map[string]any{
		"motion_id": 357,
	})
	tc.AssertSuccess(resp, err)
}
