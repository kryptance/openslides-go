package meeting

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("meeting.replace_projector_id", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":           map[string]any{"type": "integer"},
			"projector_id": map[string]any{"type": "integer"},
			"place":        map[string]any{"type": "string"},
		},
		"required": []string{"id", "projector_id", "place"},
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

		projectorID := instance["projector_id"]
		place, ok := instance["place"].(string)
		if !ok {
			return nil, fmt.Errorf("place must be a string")
		}

		// Build the projector place field name (e.g., "reference_projector_id").
		fieldName := fmt.Sprintf("%s_projector_id", place)

		fqid := event.FQID("meeting", idInt)
		e := event.Event{
			Type: event.TypeUpdate,
			FQID: fqid,
			Fields: map[string]any{
				fieldName: projectorID,
			},
		}

		return []event.Event{e}, nil
	}

	action.Register(a)
}
