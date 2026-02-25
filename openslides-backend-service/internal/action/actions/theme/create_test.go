package theme

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test"},
	})

	resp, err := tc.RequestInternal("theme.create", map[string]any{
		"name":        "test_Xcdfgee",
		"primary_500": "#111222",
		"accent_500":  "#111222",
		"warn_500":    "#222333",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("theme/1", map[string]any{
		"name":        "test_Xcdfgee",
		"primary_500": "#111222",
		"accent_500":  "#111222",
		"warn_500":    "#222333",
	})
}

func TestCreateOptFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test"},
	})

	resp, err := tc.RequestInternal("theme.create", map[string]any{
		"name":        "test_Xcdfgee",
		"primary_500": "#111222",
		"accent_500":  "#111222",
		"warn_500":    "#222333",
		"headbar":     "#333444",
		"yes":         "#333555",
		"no":          "#333666",
		"abstain":     "#333777",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("theme/1", map[string]any{
		"name":        "test_Xcdfgee",
		"primary_500": "#111222",
		"accent_500":  "#111222",
		"warn_500":    "#222333",
		"headbar":     "#333444",
		"yes":         "#333555",
		"no":          "#333666",
		"abstain":     "#333777",
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("theme.create", map[string]any{})
	tc.AssertErrorContains(err, "name")
}
