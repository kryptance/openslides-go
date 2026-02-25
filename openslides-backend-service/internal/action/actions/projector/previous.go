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
	a := action.NewBaseAction("projector.previous", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	mixin.WithSingularAction(a)

	// Show previous stable projection.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectorID := toInt(instance["id"])
		if projectorID == 0 {
			return nil, fmt.Errorf("instance has no id field")
		}

		proj, found := params.Datastore.GetChangedModel("projector", projectorID)
		if !found {
			return nil, fmt.Errorf("projector/%d not found", projectorID)
		}

		// Get current non-stable projections and move them to history.
		currentIDs, _ := proj["current_projection_ids"]
		cids, _ := toIntSlice(currentIDs)

		var events []event.Event
		for _, cid := range cids {
			projection, found := params.Datastore.GetChangedModel("projection", cid)
			if !found {
				continue
			}
			stable, _ := projection["stable"].(bool)
			if !stable {
				fields := map[string]any{
					"current_projector_id": nil,
					"history_projector_id": projectorID,
				}
				params.Datastore.ApplyChangedModel("projection", cid, fields)
				fqid := event.FQID("projection", cid)
				events = append(events, event.Event{
					Type:   event.TypeUpdate,
					FQID:   fqid,
					Fields: fields,
				})
			}
		}

		return events, nil
	}

	action.Register(a)
}
