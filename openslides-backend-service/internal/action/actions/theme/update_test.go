package theme

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test"},
		"theme/1":        {"name": "old", "organization_id": 1},
	})

	resp, err := tc.RequestInternal("theme.update", map[string]any{
		"id":          1,
		"name":        "test",
		"primary_500": "#121212",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("theme/1", map[string]any{
		"name":        "test",
		"primary_500": "#121212",
	})
}

func TestUpdateOptFieldsCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test"},
		"theme/1":        {"name": "old", "organization_id": 1},
	})

	resp, err := tc.RequestInternal("theme.update", map[string]any{
		"id":      1,
		"name":    "test",
		"headbar": "#333444",
		"yes":     "#333555",
		"no":      "#333666",
		"abstain": "#333777",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("theme/1", map[string]any{
		"name":    "test",
		"headbar": "#333444",
		"yes":     "#333555",
		"no":      "#333666",
		"abstain": "#333777",
	})
}
