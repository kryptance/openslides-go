package projector_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {"projector_message_ids": []any{2}},
		"projector_message/2": {
			"meeting_id": 1,
			"message":    "test1",
		},
	})

	resp, err := tc.Request("projector_message.update", map[string]any{
		"id":      2,
		"message": "geredegerede",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_message/2")
	if model["message"] != "geredegerede" {
		t.Errorf("expected message 'geredegerede', got %v", model["message"])
	}
}

func TestUpdateMissingID(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("projector_message.update", map[string]any{
		"message": "geredegerede",
	})
	tc.AssertErrorContains(err, "id")
}
