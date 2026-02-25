package motion_editor

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/357": {
			"title":      "title_YIDYXmKj",
			"meeting_id": 1,
		},
		"user/78": {
			"username":         "username_loetzbfg",
			"meeting_ids":      []any{1},
			"meeting_user_ids": []any{78},
		},
		"meeting_user/78": {
			"meeting_id": 1,
			"user_id":    78,
		},
		"meeting_user/79": {
			"meeting_id": 1,
			"user_id":    78,
		},
		"motion_editor/1": {
			"meeting_user_id": 78,
			"meeting_id":      1,
			"motion_id":       357,
		},
	})
	resp, err := tc.Request("motion_editor.update", map[string]any{
		"id":              1,
		"meeting_user_id": 79,
	})
	tc.AssertSuccess(resp, err)
}

func TestUpdateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_editor.update", map[string]any{})
	tc.AssertError(resp, err)
}
