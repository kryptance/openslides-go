package meeting

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("meeting.set_logo", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":           map[string]any{"type": "integer"},
			"logo":         map[string]any{"type": "string"},
			"mediafile_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "logo", "mediafile_id"},
	}

	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
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

		logo, ok := instance["logo"].(string)
		if !ok {
			return nil, fmt.Errorf("logo must be a string")
		}

		mediafileID := instance["mediafile_id"]

		// Build the logo field name (e.g., "logo_web_header_id", "logo_pdf_header_l_id").
		fieldName := fmt.Sprintf("logo_%s_id", logo)

		fqid := event.FQID("meeting", idInt)
		e := event.Event{
			Type: event.TypeUpdate,
			FQID: fqid,
			Fields: map[string]any{
				fieldName: mediafileID,
			},
		}

		return []event.Event{e}, nil
	}

	action.Register(a)
}
