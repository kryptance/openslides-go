package organization_tag

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1":     {"name": "test", "organization_tag_ids": []any{1}},
		"organization_tag/1": {"name": "test", "color": "#000000", "organization_id": 1},
	})

	resp, err := tc.RequestInternal("organization_tag.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("organization_tag/1")
}
