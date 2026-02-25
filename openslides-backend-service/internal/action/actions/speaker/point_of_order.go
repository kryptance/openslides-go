package speaker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("speaker.point_of_order", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanBeSpeaker

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"list_of_speakers_id":        map[string]any{"type": "integer"},
			"note":                        map[string]any{"type": "string"},
			"point_of_order_category_id":  map[string]any{"type": "integer"},
		},
		"required": []string{"list_of_speakers_id"},
	}

	mixin.WithSingularAction(a)

	action.Register(a)
}
