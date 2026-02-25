package committee

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCommitteeJsonUploadSimple(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.json_upload", map[string]any{
		"data": []any{
			map[string]any{
				"name": "test",
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeJsonUploadMultiple(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.json_upload", map[string]any{
		"data": []any{
			map[string]any{
				"name": "committee A",
			},
			map[string]any{
				"name": "committee B",
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeJsonUploadWithDescription(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.json_upload", map[string]any{
		"data": []any{
			map[string]any{
				"name":        "test",
				"description": "desc",
			},
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeJsonUploadMissingData(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.json_upload", map[string]any{})
	tc.AssertError(resp, err)
}
