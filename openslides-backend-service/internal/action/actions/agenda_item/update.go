package agenda_item

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("agenda_item.update", model.MustGet("agenda_item"))
	a.Permission = perm.AgendaItemCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":          map[string]any{"type": "integer"},
			"item_number": map[string]any{"type": "string"},
			"comment":     map[string]any{"type": "string"},
			"type":        map[string]any{"type": "integer"},
			"parent_id":   map[string]any{"type": "integer"},
			"weight":      map[string]any{"type": "integer"},
			"duration":    map[string]any{"type": "integer"},
			"closed":      map[string]any{"type": "boolean"},
			"tag_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
