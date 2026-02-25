package projector_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"projector_message_ids": []any{2}},
		"projector_message/2": {
			"meeting_id": 1,
			"message":    "test1",
		},
	})

	resp, err := tc.Request("projector_message.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("projector_message/2")
}
