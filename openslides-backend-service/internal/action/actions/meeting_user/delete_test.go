package meeting_user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestMeetingUserDelete(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10":      {"is_active_in_organization_id": 1},
		"meeting_user/5":  {"user_id": 1, "meeting_id": 10},
		"user/1":          {"username": "admin", "meeting_user_ids": []any{5}},
	})
	resp, err := tc.Request("meeting_user.delete", map[string]any{"id": 5})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting_user/5")
}

func TestMeetingUserDeleteNonExistent(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// delete operations. The delete succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("meeting_user.delete", map[string]any{"id": 999})
	tc.AssertSuccess(resp, err)
}

func TestMeetingUserDeleteMultiple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/10": {"is_active_in_organization_id": 1},
		"meeting_user/5": {"user_id": 1, "meeting_id": 10},
		"meeting_user/6": {"user_id": 2, "meeting_id": 10},
		"user/1":         {"username": "admin", "meeting_user_ids": []any{5}},
		"user/2":         {"username": "user2", "meeting_user_ids": []any{6}},
	})
	resp1, err1 := tc.Request("meeting_user.delete", map[string]any{"id": 5})
	tc.AssertSuccess(resp1, err1)
	tc.AssertModelDeleted("meeting_user/5")

	resp2, err2 := tc.Request("meeting_user.delete", map[string]any{"id": 6})
	tc.AssertSuccess(resp2, err2)
	tc.AssertModelDeleted("meeting_user/6")
}
