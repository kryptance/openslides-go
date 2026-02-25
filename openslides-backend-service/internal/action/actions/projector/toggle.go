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
	a := action.NewBaseAction("projector.toggle", model.MustGet("projector"))
	a.Permission = perm.ProjectorCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":                map[string]any{"type": "integer"},
			"content_object_id": map[string]any{"type": "string"},
			"meeting_id":        map[string]any{"type": "integer"},
			"stable":            map[string]any{"type": "boolean"},
			"type":              map[string]any{"type": "string"},
		},
		"required": []string{"id", "content_object_id", "meeting_id"},
	}

	mixin.WithSingularAction(a)

	// Toggle projection: if the content is currently projected, remove it; otherwise add it.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		projectorID := toInt(instance["id"])
		if projectorID == 0 {
			return nil, fmt.Errorf("instance has no id field")
		}

		contentObjectID, _ := instance["content_object_id"].(string)
		meetingID := toInt(instance["meeting_id"])
		stable, _ := instance["stable"].(bool)
		projType, _ := instance["type"].(string)

		proj, found := params.Datastore.GetChangedModel("projector", projectorID)
		if !found {
			return nil, fmt.Errorf("projector/%d not found", projectorID)
		}

		// Check if the content is already projected.
		currentIDs, _ := proj["current_projection_ids"]
		cids, _ := toIntSlice(currentIDs)

		for _, cid := range cids {
			projection, found := params.Datastore.GetChangedModel("projection", cid)
			if !found {
				continue
			}
			existingCOID, _ := projection["content_object_id"].(string)
			if existingCOID == contentObjectID {
				// Already projected: remove it.
				fqid := event.FQID("projection", cid)
				return []event.Event{{
					Type: event.TypeDelete,
					FQID: fqid,
				}}, nil
			}
		}

		// Not projected: create a new projection.
		projectionIDs, err := params.Datastore.ReserveIDs("projection", 1)
		if err != nil {
			return nil, fmt.Errorf("reserve projection id: %w", err)
		}
		projectionID := projectionIDs[0]

		projFields := map[string]any{
			"id":                   projectionID,
			"meeting_id":           meetingID,
			"content_object_id":    contentObjectID,
			"current_projector_id": projectorID,
			"stable":               stable,
		}
		if projType != "" {
			projFields["type"] = projType
		}

		params.Datastore.ApplyChangedModel("projection", projectionID, projFields)

		projFQID := event.FQID("projection", projectionID)
		return []event.Event{{
			Type:   event.TypeCreate,
			FQID:   projFQID,
			Fields: projFields,
		}}, nil
	}

	action.Register(a)
}
