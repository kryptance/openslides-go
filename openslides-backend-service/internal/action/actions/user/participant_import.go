package user

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("user.participant.import", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":         map[string]any{"type": "integer"},
			"meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "meeting_id"},
	}

	importAction := mixin.NewBaseImportAction("user.participant.json_upload")
	importAction.AddLookup("username", mixin.NewLookup("user", "username"))

	mixin.WithBaseImport(a, importAction)
	mixin.WithSingularAction(a)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		_ = id

		// TODO: Load the import preview data (action_worker) by id.
		// TODO: For each row in the import data:
		//   - If user is new: create user account and meeting_user.
		//   - If user exists: create/update meeting_user for the specified meeting.
		//   - Assign groups, structure levels, vote weight from the row data.
		//   - Skip rows in "error" state.
		// TODO: Generate create/update events for users and meeting_users.
		// TODO: Mark the action_worker as completed.

		return nil, fmt.Errorf("user.participant.import: not yet implemented")
	}

	action.Register(a)
}
