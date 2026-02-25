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
	a := action.NewBaseAction("projector.project_preview", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	mixin.WithSingularAction(a)

	// Project the first preview item onto the projector.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectorID := toInt(instance["id"])
		if projectorID == 0 {
			return nil, fmt.Errorf("instance has no id field")
		}

		proj, found := params.Datastore.GetChangedModel("projector", projectorID)
		if !found {
			return nil, fmt.Errorf("projector/%d not found", projectorID)
		}

		previewIDs, ok := proj["preview_projection_ids"]
		if !ok || previewIDs == nil {
			return nil, fmt.Errorf("no preview projections on projector/%d", projectorID)
		}

		pids, ok := toIntSlice(previewIDs)
		if !ok || len(pids) == 0 {
			return nil, fmt.Errorf("no preview projections on projector/%d", projectorID)
		}

		// Move the first preview projection to current.
		firstID := pids[0]
		fields := map[string]any{
			"current_projector_id": projectorID,
			"preview_projector_id": nil,
		}
		params.Datastore.ApplyChangedModel("projection", firstID, fields)

		fqid := event.FQID("projection", firstID)
		return []event.Event{{
			Type:   event.TypeUpdate,
			FQID:   fqid,
			Fields: fields,
		}}, nil
	}

	action.Register(a)
}
