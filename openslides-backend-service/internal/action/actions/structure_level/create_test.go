package structure_level

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateEmptyData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.Request("structure_level.create", map[string]any{})
	tc.AssertErrorContains(err, "meeting_id")
}

func TestCreateRequiredFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("structure_level.create", map[string]any{
		"name":       "test",
		"meeting_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("structure_level/1", map[string]any{
		"name":       "test",
		"meeting_id": 1,
	})
}

func TestCreateAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.Request("structure_level.create", map[string]any{
		"name":         "test",
		"meeting_id":   1,
		"color":        "#abf257",
		"default_time": 600,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("structure_level/1", map[string]any{
		"name":         "test",
		"meeting_id":   1,
		"color":        "#abf257",
		"default_time": 600,
	})
}
