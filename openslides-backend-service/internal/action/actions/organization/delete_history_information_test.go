package organization

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteHistoryInformationCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name": "Orga",
		},
	})
	resp, err := tc.Request("organization.delete_history_information", map[string]any{
		"id": 1,
	})
	tc.AssertSuccess(resp, err)
}

func TestDeleteHistoryInformationMissingId(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("organization.delete_history_information", map[string]any{})
	tc.AssertError(resp, err)
}
