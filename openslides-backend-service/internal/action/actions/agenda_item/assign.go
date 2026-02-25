package agenda_item

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("agenda_item.assign", model.MustGet("agenda_item"))
	a.Permission = perm.AgendaItemCanManage
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
			"ids":        map[string]any{"type": "array", "items": map[string]any{"type": "integer"}},
			"parent_id":  map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "ids"},
	}

	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Assign all agenda items in ids to the given parent_id.
		// For each item: set parent_id, recalculate tree weights.
		return instance, nil
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		// TODO: Generate update events for all affected agenda items.
		return nil, nil
	}

	action.Register(a)
}
