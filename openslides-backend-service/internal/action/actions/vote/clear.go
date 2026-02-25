package vote

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
)

func init() {
	a := action.NewUpdateAction("vote.clear", model.MustGet("vote"))
	a.ActionType = action.ActionTypeBackendInternal

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"poll_id": map[string]any{"type": "integer"},
		},
		"required": []string{"poll_id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Look up all votes belonging to the given poll_id
		// (via poll -> option_ids -> vote_ids) and delete them.
		// This would typically execute vote.delete sub-actions for each vote.

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
