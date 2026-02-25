package mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"mediafile/111": {
			"title":    "title_srtgb123",
			"owner_id": "meeting/1",
		},
	})

	resp, err := tc.Request("mediafile.update", map[string]any{
		"id":    111,
		"title": "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/111", map[string]any{
		"title": "title_Xcdfgee",
	})
}

func TestUpdateToken(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"mediafile/7": {
			"token":    "token_1",
			"owner_id": "organization/1",
		},
	})

	resp, err := tc.Request("mediafile.update", map[string]any{
		"id":    7,
		"token": "token_1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/7", map[string]any{
		"token": "token_1",
	})
}
