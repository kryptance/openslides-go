package topic

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// TestJsonUploadBasic tests basic json_upload with valid data.
// Note: The full json_upload workflow (duplicate detection, import_preview creation,
// statistics, etc.) is not yet fully implemented in Go. These tests verify the basic
// action structure and schema validation.
func TestJsonUploadBasic(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.json_upload", map[string]any{
		"meeting_id": 1,
		"data": []any{
			map[string]any{"title": "test"},
		},
	})
	tc.AssertSuccess(resp, err)
}

// TestJsonUploadMissingMeetingID tests that missing meeting_id causes an error.
func TestJsonUploadMissingMeetingID(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("topic.json_upload", map[string]any{
		"data": []any{
			map[string]any{"title": "test"},
		},
	})
	tc.AssertErrorContains(err, "meeting_id")
}

// TestJsonUploadMissingData tests that missing data causes an error.
func TestJsonUploadMissingData(t *testing.T) {
	tc := testutil.New(t)

	_, err := tc.RequestInternal("topic.json_upload", map[string]any{
		"meeting_id": 1,
	})
	tc.AssertErrorContains(err, "data")
}

// TestJsonUploadMultipleRows tests uploading multiple topic rows.
func TestJsonUploadMultipleRows(t *testing.T) {
	tc := testutil.New(t)
	tc.CreateMeeting()

	resp, err := tc.RequestInternal("topic.json_upload", map[string]any{
		"meeting_id": 1,
		"data": []any{
			map[string]any{"title": "test1"},
			map[string]any{"title": "test2"},
		},
	})
	tc.AssertSuccess(resp, err)
}
