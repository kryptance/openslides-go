package option

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("option.update", model.MustGet("option"))
	a.Permission = perm.PollCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":      map[string]any{"type": "integer"},
			"text":    map[string]any{"type": "string"},
			"weight":  map[string]any{"type": "integer"},
			"yes":     map[string]any{"type": "string"},
			"no":      map[string]any{"type": "string"},
			"abstain": map[string]any{"type": "string"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
