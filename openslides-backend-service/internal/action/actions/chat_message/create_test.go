package chat_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{1},
		},
		"chat_group/2":   {"meeting_id": 1, "write_group_ids": []any{3}},
		"group/3":        {"meeting_id": 1, "meeting_user_ids": []any{1}},
		"user/1":         {"meeting_user_ids": []any{1}},
		"meeting_user/1": {"meeting_id": 1, "user_id": 1, "group_ids": []any{3}},
	})

	resp, err := tc.Request("chat_message.create", map[string]any{
		"chat_group_id": 2,
		"content":       "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("chat_message/1", map[string]any{
		"content":       "test",
		"chat_group_id": 2,
	})
}
