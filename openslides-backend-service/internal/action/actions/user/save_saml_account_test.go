package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// save_saml_account is not yet implemented.
func TestSaveSamlAccountNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.save_saml_account", map[string]any{
		"username": "saml_user",
		"saml_id":  "saml123",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}

func TestSaveSamlAccountEmptySamlId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.save_saml_account", map[string]any{
		"username": "saml_user",
		"saml_id":  "",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "saml_id")
}
