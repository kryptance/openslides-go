package agenda_item

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
	a := action.NewBaseAction("agenda_item.numbering", model.MustGet("agenda_item"))
	a.Permission = perm.AgendaItemCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"meeting_id": map[string]any{"type": "integer"},
		},
		"required": []string{"meeting_id"},
	}

	mixin.WithSingularAction(a)

	// Auto-number all agenda items in the meeting.
	a.CreateEvents = func(ctx context.Context, params *action.ActionParams, instance action.Instance) ([]event.Event, error) {
		meetingID := toInt(instance["meeting_id"])
		if meetingID == 0 {
			return nil, fmt.Errorf("meeting_id is required")
		}

		// Look up meeting to find its agenda_item_ids.
		meeting, found := params.Datastore.GetChangedModel("meeting", meetingID)
		if !found {
			return nil, fmt.Errorf("meeting/%d not found", meetingID)
		}

		agendaItemIDs, ok := meeting["agenda_item_ids"]
		if !ok || agendaItemIDs == nil {
			return nil, nil
		}

		aids, ok := toIntSlice(agendaItemIDs)
		if !ok || len(aids) == 0 {
			return nil, nil
		}

		// Assign sequential item numbers.
		var events []event.Event
		for i, aid := range aids {
			itemNumber := fmt.Sprintf("%d", i+1)
			fields := map[string]any{"item_number": itemNumber}
			fqid := event.FQID("agenda_item", aid)
			events = append(events, event.Event{
				Type:   event.TypeUpdate,
				FQID:   fqid,
				Fields: fields,
			})
			params.Datastore.ApplyChangedModel("agenda_item", aid, fields)
		}

		return events, nil
	}

	action.Register(a)
}
