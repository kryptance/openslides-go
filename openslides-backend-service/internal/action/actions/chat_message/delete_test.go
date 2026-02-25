package chat_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1":       {"is_active_in_organization_id": 1, "meeting_user_ids": []any{5}},
		"chat_group/11":   {"meeting_id": 1, "name": "test"},
		"chat_message/101": {"meeting_id": 1, "meeting_user_id": 5, "chat_group_id": 11},
		"meeting_user/5":  {"meeting_id": 1, "user_id": 3, "chat_message_ids": []any{101}},
		"user/3":          {"username": "username_xx", "meeting_user_ids": []any{5}},
	})

	resp, err := tc.Request("chat_message.delete", map[string]any{"id": 101})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("chat_message/101")
}
