package chat_group

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

func init() {
	a := action.NewBaseAction("chat_group.clear", model.MustGet("chat_group"))
	a.Permission = perm.ChatCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	// CreateEvents deletes all chat_message instances belonging to this chat_group.
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

		// Look up the chat_group to find its chat_message_ids.
		chatGroup, found := params.Datastore.GetChangedModel("chat_group", idInt)
		if !found {
			return nil, fmt.Errorf("chat_group/%d not found", idInt)
		}

		messageIDs, ok := chatGroup["chat_message_ids"]
		if !ok || messageIDs == nil {
			return nil, nil
		}

		msgSlice, ok := messageIDs.([]any)
		if !ok {
			return nil, nil
		}

		var events []event.Event
		for _, mid := range msgSlice {
			var msgID int
			switch v := mid.(type) {
			case float64:
				msgID = int(v)
			case int:
				msgID = v
			default:
				continue
			}

			// Mark as deleted in the datastore so AssertModelDeleted works.
			params.Datastore.ApplyToBeDeleted("chat_message", msgID)

			fqid := event.FQID("chat_message", msgID)
			events = append(events, event.Event{
				Type: event.TypeDelete,
				FQID: fqid,
			})
		}

		return events, nil
	}

	action.Register(a)
}
