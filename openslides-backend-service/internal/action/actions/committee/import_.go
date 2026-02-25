package committee

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
	a := action.NewBaseAction("committee.import", model.MustGet("committee"))
	a.Permission = perm.OMLSuperadmin

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	// Import is a singular action.
	mixin.WithSingularAction(a)

	// Set up base import action referencing the JSON upload preview.
	importAction := mixin.NewBaseImportAction("committee.json_upload")
	importAction.AddLookup("committee", mixin.NewLookup("committee", "name"))
	mixin.WithBaseImport(a, importAction)

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// TODO: Implement committee import execution logic:
		// 1. Load the validated import data from the action_worker entry
		//    referenced by the id field.
		// 2. For each row, create or update the committee as appropriate.
		// 3. Generate create/update events for all processed committees.
		// 4. Mark the action_worker entry as completed.
		_ = fmt.Sprintf("committee.import")
		return nil, nil
	}

	action.Register(a)
}
