package topic

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("topic.update", model.MustGet("topic"))
	a.Permission = perm.AgendaItemCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":    map[string]any{"type": "integer"},
			"title": map[string]any{"type": "string", "minLength": 1},
			"text":  map[string]any{"type": "string"},
			"attachment_mediafile_ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"id"},
	}

	// TODO: mixin.WithAttachment(a)

	action.Register(a)
}
