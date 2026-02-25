package motion_supporter

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"title":        "motion_1",
			"meeting_id":   1,
			"state_id":     1,
			"supporter_ids": []any{},
		},
		"motion_state/1": {
			"name":          "state_1",
			"allow_support": true,
			"motion_ids":    []any{1},
			"meeting_id":    1,
		},
		"meeting_user/1": {
			"meeting_id": 1,
			"user_id":    1,
		},
	})
	resp, err := tc.Request("motion_supporter.create", map[string]any{
		"motion_id":       1,
		"meeting_user_id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_supporter.create", map[string]any{})
	tc.AssertError(resp, err)
}

func TestCreateWithMultipleSupporters(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/1": {
			"title":        "motion_1",
			"meeting_id":   1,
			"state_id":     1,
			"supporter_ids": []any{},
		},
		"motion_state/1": {
			"name":          "state_1",
			"allow_support": true,
			"motion_ids":    []any{1},
			"meeting_id":    1,
		},
		"user/78": {
			"username":    "username_bob",
			"meeting_ids": []any{1},
		},
		"meeting_user/78": {
			"meeting_id": 1,
			"user_id":    78,
		},
	})
	resp, err := tc.Request("motion_supporter.create", map[string]any{
		"motion_id":       1,
		"meeting_user_id": 78,
	})
	tc.AssertSuccess(resp, err)
}
