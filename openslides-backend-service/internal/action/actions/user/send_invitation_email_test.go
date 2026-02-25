package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// send_invitation_email generates no datastore events (side-effect is email).
func TestSendInvitationEmailNoEvents(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "email": "test@example.com"},
		"meeting/1": {
			"name":                         "Test Meeting",
			"committee_id":                 1,
			"is_active_in_organization_id": 1,
		},
		"committee/1": {},
	})
	resp, err := tc.Request("user.send_invitation_email", map[string]any{
		"id":         2,
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertEventCount(resp, 0)
}
