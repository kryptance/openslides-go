package organization

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("organization.set_theme", model.MustGet("organization"))
	a.Permission = perm.OMLSuperadmin

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":       map[string]any{"type": "integer"},
			"theme_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "theme_id"},
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

		themeID := instance["theme_id"]

		fqid := event.FQID("organization", idInt)
		e := event.Event{
			Type: event.TypeUpdate,
			FQID: fqid,
			Fields: map[string]any{
				"theme_id": themeID,
			},
		}

		return []event.Event{e}, nil
	}

	action.Register(a)
}
