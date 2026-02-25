package poll_candidate_list

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("poll_candidate_list.delete", model.MustGet("poll_candidate_list"))
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
