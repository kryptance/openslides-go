package mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"mediafile/111": {
			"title":    "title_srtgb123",
			"owner_id": "meeting/1",
		},
	})

	resp, err := tc.Request("mediafile.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("mediafile/111")
}

func TestDeleteOrganizationMediafile(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"mediafile/112": {
			"title":       "title_srtgb123",
			"is_directory": true,
			"child_ids":   []any{110, 113},
			"owner_id":    "organization/1",
		},
		"mediafile/110": {
			"title":     "title_ghjeu212",
			"parent_id": 112,
			"owner_id":  "organization/1",
		},
		"mediafile/113": {
			"title":     "title_del2",
			"parent_id": 112,
			"owner_id":  "organization/1",
		},
	})

	resp, err := tc.Request("mediafile.delete", map[string]any{"id": 112})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("mediafile/112")
}
