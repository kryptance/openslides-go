package mediafile

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUploadSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("mediafile.upload", map[string]any{
		"title":    "title_xXRGTLAJ",
		"owner_id": "meeting/1",
		"filename": "fn_jumbo.txt",
		"file":     "dGVzdHRlc3R0ZXN0",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/1", map[string]any{
		"title":    "title_xXRGTLAJ",
		"owner_id": "meeting/1",
		"filename": "fn_jumbo.txt",
	})
}

func TestUploadOrganization(t *testing.T) {
	tc := testutil.New(t)

	resp, err := tc.Request("mediafile.upload", map[string]any{
		"title":    "title_xXRGTLAJ",
		"owner_id": "organization/1",
		"filename": "fn_jumbo.txt",
		"file":     "dGVzdHRlc3R0ZXN0",
		"token":    "web_logo",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("mediafile/1", map[string]any{
		"title":    "title_xXRGTLAJ",
		"owner_id": "organization/1",
	})
}

func TestUploadWithParent(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"mediafile/10": {
			"title":        "title_CgKPfByo",
			"is_directory":  true,
			"owner_id":     "meeting/1",
		},
	})

	resp, err := tc.Request("mediafile.upload", map[string]any{
		"title":     "title_xXRGTLAJ",
		"owner_id":  "meeting/1",
		"filename":  "fn_jumbo.txt",
		"file":      "dGVzdHRlc3R0ZXN0",
		"parent_id": 10,
	})
	tc.AssertSuccess(resp, err)
	// SetModels with mediafile/10 sets nextID["mediafile"]=10, so the next create gets mediafile/11.
	tc.AssertModelExists("mediafile/11", map[string]any{
		"title":     "title_xXRGTLAJ",
		"owner_id":  "meeting/1",
		"filename":  "fn_jumbo.txt",
		"parent_id": 10,
	})
}
