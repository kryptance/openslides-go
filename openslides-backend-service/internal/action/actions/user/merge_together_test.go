package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// merge_together is not yet implemented.
func TestMergeTogetherNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "target_user"},
		"user/2": {"username": "source_user"},
	})
	resp, err := tc.Request("user.merge_together", map[string]any{
		"id":       1,
		"user_ids": []any{2},
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
