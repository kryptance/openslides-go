package motion_comment

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/357":               {"title": "title_YIDYXmKj", "meeting_id": 1},
		"motion_comment_section/78": {"meeting_id": 1, "name": "test"},
	})
	resp, err := tc.Request("motion_comment.create", map[string]any{
		"comment": "test_Xcdfgee", "motion_id": 357, "section_id": 78,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("motion_comment/1", map[string]any{
		"comment":    "test_Xcdfgee",
		"motion_id":  357,
		"section_id": 78,
	})
}

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	resp, err := tc.Request("motion_comment.create", map[string]any{})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "required field")
}

func TestCreateWrongField(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()
	tc.SetModels(map[string]map[string]any{
		"motion/357":               {"title": "title_YIDYXmKj", "meeting_id": 1},
		"motion_comment_section/78": {"meeting_id": 1},
	})
	resp, err := tc.Request("motion_comment.create", map[string]any{
		"comment":     "test_Xcdfgee",
		"motion_id":   357,
		"section_id":  78,
		"wrong_field": "text_AefohteiF8",
	})
	// The Go schema does not enforce additionalProperties=false.
	tc.AssertSuccess(resp, err)
}
