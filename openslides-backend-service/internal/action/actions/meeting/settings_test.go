package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSettingsGroupIds(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"motion_poll_default_group_ids":  []any{1},
			"is_active_in_organization_id":   1,
			"language":                       "en",
			"committee_id":                   1,
		},
		"group/1": {
			"used_as_motion_poll_default_id": 1,
		},
		"group/2": {
			"name":                          "2",
			"used_as_motion_poll_default_id": nil,
		},
		"group/3": {
			"used_as_motion_poll_default_id": nil,
		},
	})
	resp, err := tc.Request("meeting.update", map[string]any{
		"id":                            1,
		"motion_poll_default_group_ids": []any{2, 3},
	})
	tc.AssertSuccess(resp, err)
}
