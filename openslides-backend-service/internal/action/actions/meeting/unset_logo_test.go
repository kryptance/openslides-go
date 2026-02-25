package meeting

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUnsetLogo(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"logo_pdf_header_l_id":          7,
			"logo_pdf_header_r_id":          7,
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
			"meeting_id":                                222,
			"mediafile_id":                              17,
			"used_as_logo_pdf_header_l_in_meeting_id":   222,
			"used_as_logo_pdf_header_r_in_meeting_id":   222,
		},
	})
	resp, err := tc.Request("meeting.unset_logo", map[string]any{
		"id":   222,
		"logo": "pdf_header_l",
	})
	tc.AssertSuccess(resp, err)
}

func TestUnsetLogoWebHeader(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"logo_web_header_id":            7,
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
			"used_as_logo_web_header_in_meeting_id":   222,
		},
	})
	resp, err := tc.Request("meeting.unset_logo", map[string]any{
		"id":   222,
		"logo": "web_header",
	})
	tc.AssertSuccess(resp, err)
}

func TestUnsetLogoMissingFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"meeting/222": {
			"is_active_in_organization_id": 1,
			"committee_id":                1,
		},
	})
	resp, err := tc.Request("meeting.unset_logo", map[string]any{
		"id": 222,
	})
	tc.AssertError(resp, err)
}
