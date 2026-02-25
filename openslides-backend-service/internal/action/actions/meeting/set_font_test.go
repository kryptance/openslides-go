package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestSetFontCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"meeting_mediafile_ids":         []any{7},
			"is_active_in_organization_id":  1,
			"committee_id":                  1,
		},
		"mediafile/17": {
			"is_directory":          false,
			"mimetype":              "font/woff",
			"owner_id":             "meeting/222",
			"meeting_mediafile_ids": []any{7},
		},
		"meeting_mediafile/7": {
			"meeting_id":   222,
			"mediafile_id": 17,
		},
	})
	resp, err := tc.Request("meeting.set_font", map[string]any{
		"id":           222,
		"mediafile_id": 17,
		"font":         "bold",
	})
	tc.AssertSuccess(resp, err)
}

func TestSetFontMissingFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"is_active_in_organization_id": 1,
			"committee_id":                1,
		},
	})
	resp, err := tc.Request("meeting.set_font", map[string]any{
		"id": 222,
	})
	tc.AssertError(resp, err)
}
