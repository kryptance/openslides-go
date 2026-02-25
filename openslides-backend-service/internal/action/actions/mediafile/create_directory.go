package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("mediafile.create_directory", model.MustGet("mediafile"))
	a.Permission = perm.MediafileCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title":      map[string]any{"type": "string", "minLength": 1},
			"meeting_id": map[string]any{"type": "integer"},
			"owner_id":   map[string]any{"type": "string"},
			"parent_id":  map[string]any{"type": "integer"},
		},
		"required": []string{"title"},
	}

	// TODO: Implement custom UpdateInstance to set is_directory=true
	// and validate directory name uniqueness.

	action.Register(a)
}
