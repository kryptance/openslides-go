package gender

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestDeleteCorrectly(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"gender_ids": []any{1, 5, 6}},
		"gender/1":       {"name": "male", "organization_id": 1},
		"gender/5":       {"name": "fairy", "organization_id": 1},
		"gender/6":       {"name": "dragon", "organization_id": 1},
	})

	resp, err := tc.RequestInternal("gender.delete", map[string]any{"id": 5})
	tc.AssertSuccess(resp, err)
	tc.AssertModelDeleted("gender/5")
	tc.AssertModelExists("gender/1", map[string]any{"name": "male"})
	tc.AssertModelExists("gender/6", map[string]any{"name": "dragon"})
}
