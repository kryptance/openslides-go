package gender

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreate(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":       "test_organization1",
			"gender_ids": []any{},
		},
	})

	resp, err := tc.RequestInternal("gender.create", map[string]any{
		"name": "female",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("gender/1", map[string]any{
		"name": "female",
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("gender.create", map[string]any{})
	tc.AssertErrorContains(err, "name")
}
