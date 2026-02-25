package gender

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateCorrectly(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "test_organization1"},
		"gender/5":       {"name": "dragon", "organization_id": 1},
	})

	resp, err := tc.RequestInternal("gender.update", map[string]any{
		"id":   5,
		"name": "gender_testname_updated",
	})
	tc.AssertSuccess(resp, err)
	model := tc.GetModel("gender/5")
	if model["name"] != "gender_testname_updated" {
		t.Errorf("expected name 'gender_testname_updated', got %v", model["name"])
	}
}

func TestUpdateEmpty(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("gender.update", map[string]any{})
	tc.AssertErrorContains(err, "id")
}
