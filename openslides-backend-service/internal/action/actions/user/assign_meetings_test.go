package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// assign_meetings is not yet implemented.
func TestAssignMeetingsNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "admin"},
		"meeting/1": {
			"is_active_in_organization_id": 1,
			"committee_id":                 1,
		},
		"committee/1": {"meeting_ids": []any{1}},
	})
	resp, err := tc.Request("user.assign_meetings", map[string]any{
		"id":          1,
		"meeting_ids": []any{1},
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
