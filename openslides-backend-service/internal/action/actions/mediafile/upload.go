// Package mediafile implements mediafile actions.
//
// Mediafiles represent uploaded files and directories within meetings or the organization.
// Most actions require MediafileCanManage permission; publish requires OML.
package mediafile

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewCreateAction("mediafile.upload", model.MustGet("mediafile"))
	a.Permission = perm.MediafileCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title":      map[string]any{"type": "string", "minLength": 1},
			"meeting_id": map[string]any{"type": "integer"},
			"owner_id":   map[string]any{"type": "string"},
			"filename":   map[string]any{"type": "string"},
			"file":       map[string]any{"type": "string"},
			"parent_id":  map[string]any{"type": "integer"},
			"token":      map[string]any{"type": "string"},
		},
		"required": []string{"title"},
	}

	// TODO: Implement custom UpdateInstance to handle file upload,
	// mimetype detection, and PDF page counting.

	action.Register(a)
}
