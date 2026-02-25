package chat_group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"enable_chat": true},
		"committee/2":    {"meeting_ids": []any{1}},
		"meeting/1":      {"is_active_in_organization_id": 1, "committee_id": 2},
		"chat_group/1": {
			"meeting_id":      1,
			"name":            "redekreis1",
			"read_group_ids":  []any{1},
			"write_group_ids": []any{2},
		},
		"group/1": {"meeting_id": 1, "read_chat_group_ids": []any{1}},
		"group/2": {"meeting_id": 1, "write_chat_group_ids": []any{1}},
		"group/3": {"meeting_id": 1},
	})

	resp, err := tc.Request("chat_group.update", map[string]any{
		"id":              1,
		"name":            "test",
		"read_group_ids":  []any{2},
		"write_group_ids": []any{2, 3},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("chat_group/1", map[string]any{
		"name":            "test",
		"read_group_ids":  []any{2},
		"write_group_ids": []any{2, 3},
	})
}
