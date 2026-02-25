package group

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"group_ids": []any{1, 2, 3, 111},
		},
		"group/111": {"name": "group", "meeting_id": 1},
	})

	resp, err := tc.Request("group.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("group/111")
}
