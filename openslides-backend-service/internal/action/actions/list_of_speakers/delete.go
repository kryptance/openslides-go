package list_of_speakers

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("list_of_speakers.delete", model.MustGet("list_of_speakers"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
