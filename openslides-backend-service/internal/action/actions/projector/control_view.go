package projector

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("projector.control_view", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":    map[string]any{"type": "integer"},
			"field": map[string]any{"type": "string"},
			"direction": map[string]any{
				"type": "string",
				"enum": []string{"up", "down", "reset"},
			},
		},
		"required": []string{"id", "field", "direction"},
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

		field, _ := instance["field"].(string)
		direction, _ := instance["direction"].(string)

		// Look up current projector to get current scroll value.
		proj, found := params.Datastore.GetChangedModel("projector", idInt)
		if !found {
			return nil, fmt.Errorf("projector/%d not found", idInt)
		}

		var currentScroll int
		if s, ok := proj["scroll"]; ok && s != nil {
			switch v := s.(type) {
			case float64:
				currentScroll = int(v)
			case int:
				currentScroll = v
			}
		}

		var newScroll int
		if field == "scroll" {
			switch direction {
			case "up":
				newScroll = currentScroll + 1
			case "down":
				if currentScroll > 0 {
					newScroll = currentScroll - 1
				}
			case "reset":
				newScroll = 0
			}
		}

		fields := map[string]any{"scroll": newScroll}
		params.Datastore.ApplyChangedModel("projector", idInt, fields)

		fqid := event.FQID("projector", idInt)
		return []event.Event{{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: fields,
		}}, nil
	}

	action.Register(a)
}
