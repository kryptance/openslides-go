package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// account.import is not yet implemented.
func TestAccountImportNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.account.import", map[string]any{
		"id": 1,
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
