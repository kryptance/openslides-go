package option

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"option/111": {"meeting_id": 1, "content_object_id": "motion/1"},
		"meeting/1":  {"is_active_in_organization_id": 1},
		"motion/1":   {"option_ids": []any{111}},
	})

	resp, err := tc.RequestInternal("option.delete", map[string]any{"id": 111})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("option/111")
}
