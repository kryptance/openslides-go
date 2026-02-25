package organization_tag

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test"},
	})

	resp, err := tc.RequestInternal("organization_tag.create", map[string]any{
		"name":  "wSvQHymN",
		"color": "#eeeeee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization_tag/1", map[string]any{
		"name":  "wSvQHymN",
		"color": "#eeeeee",
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("organization_tag.create", map[string]any{})
	tc.AssertErrorContains(err, "name")
}
