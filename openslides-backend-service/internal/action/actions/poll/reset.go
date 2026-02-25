package poll

import (
	"context"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("poll.reset", model.MustGet("poll"))
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
		// Reset poll to "created" state.
		instance["state"] = "created"

		// Clear vote-related aggregated fields.
		instance["votesvalid"] = nil
		instance["votesinvalid"] = nil
		instance["votescast"] = nil
		instance["entitled_users_at_stop"] = nil

		// TODO: Delete all votes for this poll via vote.clear sub-action.
		// action.ExecuteSubAction(ctx, params, "vote.clear", []action.Instance{{"poll_id": instance["id"]}})

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
