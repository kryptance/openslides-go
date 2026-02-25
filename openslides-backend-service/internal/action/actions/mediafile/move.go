package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("mediafile.move", model.MustGet("mediafile"))
	a.Permission = perm.MediafileCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"owner_id":  map[string]any{"type": "string"},
			"parent_id": map[string]any{"type": "integer"},
			"ids": map[string]any{
				"type":  "array",
				"items": map[string]any{"type": "integer"},
			},
		},
		"required": []string{"ids", "owner_id"},
	}

	// Move is a singular action: one move operation at a time.
	mixin.WithSingularAction(a)

	// TODO: Implement custom UpdateInstance to move mediafiles to a new
	// parent directory, validating no circular references.

	action.Register(a)
}
