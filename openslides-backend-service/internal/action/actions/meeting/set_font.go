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
	a := action.NewBaseAction("meeting.set_font", model.MustGet("meeting"))
	a.Permission = perm.MeetingCanManageSettings

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id":           map[string]any{"type": "integer"},
			"font":         map[string]any{"type": "string"},
			"mediafile_id": map[string]any{"type": "integer"},
		},
		"required": []string{"id", "font", "mediafile_id"},
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

		font, ok := instance["font"].(string)
		if !ok {
			return nil, fmt.Errorf("font must be a string")
		}

		mediafileID := instance["mediafile_id"]

		// Build the font field name (e.g., "font_regular_id", "font_bold_id").
		fieldName := fmt.Sprintf("font_%s_id", font)

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
