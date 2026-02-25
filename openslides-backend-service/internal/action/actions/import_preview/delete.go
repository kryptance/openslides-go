package import_preview

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("import_preview.delete", model.MustGet("import_preview"))
	a.ActionType = action.ActionTypeBackendInternal

	action.Register(a)
}
