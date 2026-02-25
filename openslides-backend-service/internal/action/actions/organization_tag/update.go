package organization_tag

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("organization_tag.update", model.MustGet("organization_tag"))
	a.Permission = perm.OMLCanManageOrganization

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":    map[string]any{"type": "integer"},
			"name":  map[string]any{"type": "string", "minLength": 1},
			"color": map[string]any{"type": "string", "minLength": 1},
			"tagged_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
