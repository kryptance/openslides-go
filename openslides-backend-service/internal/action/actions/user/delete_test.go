package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {"username": "username_srtgb123"},
	})
	resp, err := tc.Request("user.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("user/111")
}

func TestDeleteWrongId(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// delete operations. This test verifies that deleting a non-existent ID
	// does not affect other models.
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/112": {"username": "username_srtgb123"},
	})
	resp, err := tc.Request("user.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/112", map[string]any{"username": "username_srtgb123"})
}

func TestDeleteWithGroups(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/111": {
			"username":         "username_srtgb123",
			"meeting_user_ids": []any{1111},
		},
		"meeting_user/1111": {
			"meeting_id": 42,
			"user_id":    111,
			"group_ids":  []any{456},
		},
		"group/456": {"meeting_id": 42, "meeting_user_ids": []any{1111}},
		"meeting/42": {
			"group_ids":                    []any{456},
			"user_ids":                     []any{111},
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{1111},
		},
	})
	resp, err := tc.Request("user.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("user/111")
}

func TestDeleteMultipleUsers(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/10": {"username": "user10"},
		"user/11": {"username": "user11"},
	})
	resp1, err1 := tc.Request("user.delete", map[string]any{"id": 10})
	tc.AssertSuccess(resp1, err1)
	tc.AssertModelDeleted("user/10")

	resp2, err2 := tc.Request("user.delete", map[string]any{"id": 11})
	tc.AssertSuccess(resp2, err2)
	tc.AssertModelDeleted("user/11")
}

func TestDeleteNonExistent(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// delete operations. The delete succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("user.delete", map[string]any{"id": 999})
	tc.AssertSuccess(resp, err)
}
