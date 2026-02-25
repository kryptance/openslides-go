package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// account.json_upload is not yet implemented.
func TestAccountJsonUploadNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.account.json_upload", map[string]any{
		"data": []any{
			map[string]any{
				"username": "testuser",
			},
		},
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
