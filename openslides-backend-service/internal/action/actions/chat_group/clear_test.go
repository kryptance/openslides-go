package chat_group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestClearCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"is_active_in_organization_id": 1},
		"chat_group/11": {
			"meeting_id":       1,
			"name":             "redekreis1",
			"chat_message_ids": []any{111, 112, 113},
		},
		"chat_message/111": {
			"content":       "test111",
			"chat_group_id": 11,
			"meeting_id":    1,
		},
		"chat_message/112": {
			"content":       "test222",
			"chat_group_id": 11,
			"meeting_id":    1,
		},
		"chat_message/113": {
			"content":       "test333",
			"chat_group_id": 11,
			"meeting_id":    1,
		},
	})

	resp, err := tc.Request("chat_group.clear", map[string]any{"id": 11})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("chat_message/111")
	tc.AssertModelDeleted("chat_message/112")
	tc.AssertModelDeleted("chat_message/113")
}
