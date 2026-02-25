package meeting_mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"mediafile/10": {"title": "hOi"},
	})

	resp, err := tc.Request("meeting_mediafile.create", map[string]any{
		"mediafile_id": 10,
		"meeting_id":   1,
		"is_public":    false,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("meeting_mediafile/1", map[string]any{
		"mediafile_id": 10,
		"meeting_id":   1,
		"is_public":    false,
	})
}
