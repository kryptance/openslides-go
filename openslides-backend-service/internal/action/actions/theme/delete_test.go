package theme

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"theme_id":  2,
			"theme_ids": []any{1, 2},
		},
		"theme/1": {
			"name":            "test",
			"primary_500":     "#000000",
			"organization_id": 1,
		},
		"theme/2": {
			"name":                       "OpenSlides Test",
			"organization_id":            1,
			"theme_for_organization_id":  1,
		},
	})

	resp, err := tc.RequestInternal("theme.delete", map[string]any{"id": 1})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("theme/1")
}
