package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupReplaceProjector(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"meeting/1": {
			"default_projector_motion_ids":  []any{11},
			"reference_projector_id":        20,
			"is_active_in_organization_id":  1,
		},
		"projector/11": {
			"used_as_default_projector_for_motion_in_meeting_id": 1,
		},
		"projector/20": {
			"used_as_reference_projector_meeting_id": 1,
		},
	})
}

func TestReplaceProjectorIdReplacing(t *testing.T) {
	tc := testutil.New(t)
	setupReplaceProjector(tc)
	resp, err := tc.RequestInternal("meeting.replace_projector_id", map[string]any{
		"id":           1,
		"projector_id": 11,
		"place":        "reference",
	})
	tc.AssertSuccess(resp, err)
}

func TestReplaceProjectorIdNoReplacing(t *testing.T) {
	tc := testutil.New(t)
	setupReplaceProjector(tc)
	resp, err := tc.RequestInternal("meeting.replace_projector_id", map[string]any{
		"id":           1,
		"projector_id": 12,
		"place":        "reference",
	})
	tc.AssertSuccess(resp, err)
}

func TestReplaceProjectorIdMissingFields(t *testing.T) {
	tc := testutil.New(t)
	setupReplaceProjector(tc)
	resp, err := tc.RequestInternal("meeting.replace_projector_id", map[string]any{
		"id": 1,
	})
	tc.AssertError(resp, err)
}
