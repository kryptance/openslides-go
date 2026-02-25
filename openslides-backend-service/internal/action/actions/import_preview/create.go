// Package import_preview implements import_preview actions.
//
// Import previews store the result of json_upload validations and are used
// by the import actions to execute the validated import.
package import_preview

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewCreateAction("import_preview.create", model.MustGet("import_preview"))
	a.ActionType = action.ActionTypeBackendInternal
	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"name":   map[string]any{"type": "string"},
			"state":  map[string]any{"type": "string"},
			"result": map[string]any{"type": "object"},
		},
		"required": []string{"name", "state"},
	}

	action.Register(a)
}
