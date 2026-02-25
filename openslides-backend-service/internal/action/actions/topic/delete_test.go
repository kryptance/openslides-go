package topic

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"topic/111": {"title": "title_srtgb123", "meeting_id": 1},
	})

	resp, err := tc.Request("topic.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("topic/111")
}

func TestDeleteMissingID(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("topic.delete", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
