package projector_message

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("projector_message.update", model.MustGet("projector_message"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":      map[string]any{"type": "integer"},
			"message": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
