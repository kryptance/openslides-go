package action_worker

import (
	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewDeleteAction("action_worker.delete", model.MustGet("action_worker"))
	a.ActionType = action.ActionTypeBackendInternal

	action.Register(a)
}
