package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("user.send_invitation_email", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "meeting_id"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// TODO: Fetch user data (email, name, username, default_password).
		// TODO: Fetch meeting data (name, email settings: subject template, body template, sender, reply-to).
		// TODO: Render the invitation email from the meeting's template.
		// TODO: Send the email via the configured email transport.
		// TODO: Update the user's last_email_sent timestamp if needed.

		_ = idInt

		// No datastore events; the side-effect is the email being sent.
		return nil, nil
	}

	action.Register(a)
}
