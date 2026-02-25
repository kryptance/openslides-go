package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// toggle_presence_by_number is not yet implemented and returns an error.
func TestTogglePresenceByNumberNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{34},
		},
		"user/111": {
			"username":         "username_srtgb123",
			"meeting_user_ids": []any{34},
		},
		"meeting_user/34": {"user_id": 111, "meeting_id": 1, "number": "1"},
		"committee/1":     {},
	})
	resp, err := tc.Request("user.toggle_presence_by_number", map[string]any{
		"meeting_id": 1,
		"number":     "1",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}

func TestTogglePresenceByNumberEmptyNumber(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.toggle_presence_by_number", map[string]any{
		"meeting_id": 1,
		"number":     "",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "number")
}
