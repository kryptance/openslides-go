package chat_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"is_active_in_organization_id": 1,
			"meeting_user_ids":             []any{7},
		},
		"chat_message/2": {
			"meeting_user_id": 7,
			"content":         "blablabla",
			"meeting_id":      1,
		},
		"meeting_user/7": {
			"meeting_id":       1,
			"user_id":          1,
			"chat_message_ids": []any{2},
		},
		"user/1": {"meeting_user_ids": []any{7}},
	})

	resp, err := tc.Request("chat_message.update", map[string]any{
		"id":      2,
		"content": "test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("chat_message/2", map[string]any{
		"content": "test",
	})
}
