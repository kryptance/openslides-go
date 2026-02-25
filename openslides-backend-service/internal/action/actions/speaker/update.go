package speaker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("speaker.update", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
			"speech_state": map[string]any{
				"type": "string",
			},
			"point_of_order":             map[string]any{"type": "boolean"},
			"note":                        map[string]any{"type": "string"},
			"point_of_order_category_id":  map[string]any{"type": "integer"},
			"meeting_user_id":             map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	action.Register(a)
}
