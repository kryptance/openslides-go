package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// participant.json_upload is not yet implemented.
func TestParticipantJsonUploadNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.participant.json_upload", map[string]any{
		"meeting_id": 1,
		"data": []any{
			map[string]any{
				"username": "testuser",
			},
		},
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
