package point_of_order_category

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("point_of_order_category.update", model.MustGet("point_of_order_category"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":   map[string]any{"type": "integer"},
			"text": map[string]any{"type": "string", "minLength": 1},
			"rank": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
