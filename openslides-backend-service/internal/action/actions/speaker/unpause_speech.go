package speaker

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewUpdateAction("speaker.unpause_speech", model.MustGet("speaker"))
	a.Permission = perm.ListOfSpeakersCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		id, ok := instance["id"]
		if !ok {
			return nil, fmt.Errorf("instance has no id field")
		}

		var idInt int
		switch v := id.(type) {
		case float64:
			idInt = int(v)
		case int:
			idInt = v
		default:
			return nil, fmt.Errorf("id has unexpected type %T", id)
		}

		// Look up current speaker to get pause_time.
		speaker, found := params.Datastore.GetChangedModel("speaker", idInt)
		if !found {
			return nil, fmt.Errorf("speaker/%d not found", idInt)
		}

		now := int(time.Now().Unix())

		// Calculate pause duration and add to total_pause.
		var pauseTime int
		if pt, ok := speaker["pause_time"]; ok && pt != nil {
			switch v := pt.(type) {
			case float64:
				pauseTime = int(v)
			case int:
				pauseTime = v
			}
		}

		var totalPause int
		if tp, ok := speaker["total_pause"]; ok && tp != nil {
			switch v := tp.(type) {
			case float64:
				totalPause = int(v)
			case int:
				totalPause = v
			}
		}

		if pauseTime > 0 {
			totalPause += now - pauseTime
		}

		instance["pause_time"] = nil
		instance["total_pause"] = totalPause

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}
