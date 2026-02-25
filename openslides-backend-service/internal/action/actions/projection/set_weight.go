package projection

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("projection.set_weight", model.MustGet("projection"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":     map[string]any{"type": "integer"},
			"weight": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "weight"},
	}

	action.Register(a)
}
