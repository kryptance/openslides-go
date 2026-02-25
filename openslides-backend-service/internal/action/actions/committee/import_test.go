package committee

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCommitteeImportSimple(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.import", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestCommitteeImportMissingId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("committee.import", map[string]any{})
	tc.AssertError(resp, err)
}
