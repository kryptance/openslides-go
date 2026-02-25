package mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateDirectoryCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("mediafile.create_directory", map[string]any{
		"owner_id": "meeting/1",
		"title":    "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/1", map[string]any{
		"title":    "title_Xcdfgee",
		"owner_id": "meeting/1",
	})
}

func TestCreateDirectoryOrganizationCorrect(t *testing.T) {
	tc := testutil.New(t)

	resp, err := tc.Request("mediafile.create_directory", map[string]any{
		"owner_id": "organization/1",
		"title":    "title_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/1", map[string]any{
		"title":    "title_Xcdfgee",
		"owner_id": "organization/1",
	})
}

func TestCreateDirectoryWithParent(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"mediafile/110": {
			"title":        "title_srtgb199",
			"owner_id":     "meeting/1",
			"is_directory":  true,
		},
	})

	resp, err := tc.Request("mediafile.create_directory", map[string]any{
		"owner_id":  "meeting/1",
		"title":     "title_Xcdfgee",
		"parent_id": 110,
	})
	tc.AssertSuccess(resp, err)
	// SetModels with mediafile/110 sets nextID["mediafile"]=110, so the next create gets mediafile/111.
	tc.AssertModelExists("mediafile/111", map[string]any{
		"title":     "title_Xcdfgee",
		"parent_id": 110,
		"owner_id":  "meeting/1",
	})
}
