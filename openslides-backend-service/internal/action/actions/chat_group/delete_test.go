package chat_group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"enable_chat": true},
		"committee/2":    {"meeting_ids": []any{1}},
		"meeting/1":      {"is_active_in_organization_id": 1, "committee_id": 2},
		"chat_group/1":   {"meeting_id": 1, "name": "redekreis1"},
	})

	resp, err := tc.Request("chat_group.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("chat_group/1")
}
