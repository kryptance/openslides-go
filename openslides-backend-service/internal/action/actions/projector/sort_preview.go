package projector

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
	a := action.NewBaseAction("projector.sort_preview", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"projection_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id", "projection_ids"},
	}

	mixin.WithSingularAction(a)

	// Sort preview projections by assigning sequential weights.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectionIDsRaw, ok := instance["projection_ids"]
		if !ok {
			return nil, fmt.Errorf("projection_ids field is required")
		}

		projectionIDs, ok := toIntSlice(projectionIDsRaw)
		if !ok {
			return nil, fmt.Errorf("projection_ids must be an array of integers")
		}

		var events []event.Event
		for i, pid := range projectionIDs {
			weight := i + 1
			fields := map[string]any{"weight": weight}
			fqid := event.FQID("projection", pid)
			events = append(events, event.Event{
				Type:   event.TypeUpdate,
				FQID:   fqid,
				Fields: fields,
			})
			params.Datastore.ApplyChangedModel("projection", pid, fields)
		}

		return events, nil
	}

	action.Register(a)
}
