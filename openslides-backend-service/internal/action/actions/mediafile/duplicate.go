package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("mediafile.duplicate", model.MustGet("mediafile"))
	a.Permission = perm.MediafileCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
			"target_id": map[string]any{"type": "integer"},
			"owner_id":  map[string]any{"type": "string"},
		},
		"required": []string{"ids"},
	}

	// Duplicate is a singular action.
	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to duplicate mediafiles
	// to a target directory, creating copies of file data.

	action.Register(a)
}
