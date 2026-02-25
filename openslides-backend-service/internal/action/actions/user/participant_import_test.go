package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

// participant.import is not yet implemented.
func TestParticipantImportNotImplemented(t *testing.T) {
	tc := testutil.New(t)
	resp, err := tc.Request("user.participant.import", map[string]any{
		"id":         1,
		"meeting_id": 1,
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "not yet implemented")
}
