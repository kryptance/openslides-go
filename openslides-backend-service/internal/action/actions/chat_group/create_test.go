package chat_group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"enable_chat": true},
	})

	resp, err := tc.Request("chat_group.create", map[string]any{
		"name":       "redekreis1",
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("chat_group/1", map[string]any{
		"name":       "redekreis1",
		"meeting_id": 1,
	})
}

func TestCreateOptionalFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"enable_chat": true},
	})

	resp, err := tc.Request("chat_group.create", map[string]any{
		"name":            "redekreis1",
		"meeting_id":      1,
		"read_group_ids":  []any{1},
		"write_group_ids": []any{2},
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("chat_group/1", map[string]any{
		"name":            "redekreis1",
		"meeting_id":      1,
		"read_group_ids":  []any{1},
		"write_group_ids": []any{2},
	})
}
