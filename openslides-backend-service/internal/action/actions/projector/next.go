package projector

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/action/mixin"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("projector.next", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	mixin.WithSingularAction(a)

	// Show next stable projection.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectorID := toInt(instance["id"])
		if projectorID == 0 {
			return nil, fmt.Errorf("instance has no id field")
		}

		proj, found := params.Datastore.GetChangedModel("projector", projectorID)
		if !found {
			return nil, fmt.Errorf("projector/%d not found", projectorID)
		}

		// Get history_projection_ids to find what to show next.
		historyIDs, _ := proj["history_projection_ids"]
		hids, _ := toIntSlice(historyIDs)

		if len(hids) == 0 {
			return nil, nil
		}

		// Move the first history projection to current.
		nextID := hids[0]
		fields := map[string]any{
			"current_projector_id": projectorID,
			"history_projector_id": nil,
		}
		params.Datastore.ApplyChangedModel("projection", nextID, fields)

		fqid := event.FQID("projection", nextID)
		return []event.Event{{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: fields,
		}}, nil
	}

	action.Register(a)
}
