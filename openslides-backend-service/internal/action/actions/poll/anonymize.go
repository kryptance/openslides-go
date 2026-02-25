package poll

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("poll.anonymize", model.MustGet("poll"))
	a.Permission = perm.PollCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		// TODO: Validate current state is "finished" or "published" and type is "named".

		// Mark the poll as pseudoanonymized.
		instance["is_pseudoanonymized"] = true

		// TODO: Remove user references from all votes of this poll.
		// For each vote belonging to this poll, set user_id and delegated_user_id to nil.

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
