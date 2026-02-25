package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("mediafile.update", model.MustGet("mediafile"))
	a.Permission = perm.MediafileCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":    map[string]any{"type": "integer"},
			"title": map[string]any{"type": "string", "minLength": 1},
			"token": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
