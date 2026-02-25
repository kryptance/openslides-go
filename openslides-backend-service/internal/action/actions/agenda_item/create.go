// Package agenda_item implements agenda_item actions.
//
// Agenda items represent entries in a meeting's agenda. They are attached to
// content objects (topics, motions, etc.) and support tree-based ordering.
package agenda_item

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("agenda_item.create", model.MustGet("agenda_item"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":        map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
			"item_number":       map[string]any{"type": "string"},
			"comment":           map[string]any{"type": "string"},
			"type":              map[string]any{"type": "integer"},
			"parent_id":         map[string]any{"type": "integer"},
			"weight":            map[string]any{"type": "integer"},
			"duration":          map[string]any{"type": "integer"},
			"tag_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"meeting_id", "content_object_id"},
	}

	mixin.WithWeight(a, "weight")
	mixin.WithTreeSort(a, "weight", "parent_id")

	action.Register(a)
}
