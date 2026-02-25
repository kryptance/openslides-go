// Package projection implements projection actions.
//
// Projections represent content that is displayed on a projector. They connect
// a content object to a projector with display options.
package projection

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("projection.create", model.MustGet("projection"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id":           map[string]any{"type": "integer"},
			"content_object_id":    map[string]any{"type": "string"},
			"current_projector_id": map[string]any{"type": "integer"},
			"preview_projector_id": map[string]any{"type": "integer"},
			"stable":               map[string]any{"type": "boolean"},
			"type":                 map[string]any{"type": "string"},
			"options":              map[string]any{"type": "object"},
			"weight":               map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id", "content_object_id"},
	}

	action.Register(a)
}
