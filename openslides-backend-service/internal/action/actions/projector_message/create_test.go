package projector_message

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("projector_message.create", map[string]any{
		"meeting_id": 1,
		"message":    "TEST",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("projector_message/1")
	if model["meeting_id"] != 1 {
		t.Errorf("expected meeting_id 1, got %v", model["meeting_id"])
	}
	if model["message"] != "TEST" {
		t.Errorf("expected message 'TEST', got %v", model["message"])
	}
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("projector_message.create", map[string]any{})
	tc.AssertErrorContains(err, "meeting_id")
}
