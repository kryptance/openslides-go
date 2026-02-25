package meeting_mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"mediafile_ids":         []any{10},
			"meeting_mediafile_ids": []any{2},
		},
		"mediafile/10": {
			"title":                "hOi",
			"meeting_mediafile_ids": []any{2},
			"owner_id":             "meeting/1",
		},
		"meeting_mediafile/2": {
			"meeting_id":                1,
			"access_group_ids":          []any{},
			"inherited_access_group_ids": []any{},
			"mediafile_id":              10,
			"is_public":                 true,
		},
	})

	resp, err := tc.Request("meeting_mediafile.delete", map[string]any{"id": 2})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("meeting_mediafile/2")
	tc.AssertModelExists("mediafile/10")
}
