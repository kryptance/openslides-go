package vote

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"vote/111":  {"meeting_id": 1},
		"meeting/1": {"is_active_in_organization_id": 1},
	})

	resp, err := tc.Request("vote.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("vote/111")
}
