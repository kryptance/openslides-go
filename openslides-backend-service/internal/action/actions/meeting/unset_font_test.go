package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUnsetFont(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"font_projector_h1_id":          7,
			"font_projector_h2_id":          7,
			"meeting_mediafile_ids":         []any{7},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
		"mediafile/17": {
			"is_directory":          false,
			"mimetype":              "image/png",
			"owner_id":             "meeting/222",
			"meeting_mediafile_ids": []any{7},
		},
		"meeting_mediafile/7": {
			"meeting_id":                              222,
			"mediafile_id":                            17,
			"used_as_font_projector_h1_in_meeting_id": 222,
			"used_as_font_projector_h2_in_meeting_id": 222,
		},
	})
	resp, err := tc.Request("meeting.unset_font", map[string]any{
		"id":   222,
		"font": "projector_h1",
	})
	tc.AssertSuccess(resp, err)
}

func TestUnsetFontMissingFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"is_active_in_organization_id": 1,
			"committee_id":                1,
		},
	})
	resp, err := tc.Request("meeting.unset_font", map[string]any{
		"id": 222,
	})
	tc.AssertError(resp, err)
}
