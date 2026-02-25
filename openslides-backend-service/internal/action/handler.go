package action

import (
	"context"
	"errors"
	"fmt"
	"log"

	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-backend-service/internal/datastore"
	"github.com/OpenSlides/openslides-backend-service/internal/event"
	"github.com/OpenSlides/openslides-backend-service/internal/history"
	"github.com/OpenSlides/openslides-backend-service/internal/relation"
)

const maxRetry = 3

// ActionResponse is the response for a single action in the payload.
type ActionResponse struct {
	Results []map[string]any
}

// HandleRequest processes an action request payload atomically.
func HandleRequest(ctx context.Context, payload []ActionPayloadElement, userID int, internal bool) ([]ActionResponse, error) {
	return executeWithRetry(ctx, payload, userID, internal, true)
}

// HandleSeparately processes each action in the payload non-atomically.
func HandleSeparately(ctx context.Context, payload []ActionPayloadElement, userID int) ([]ActionResponse, error) {
	responses := make([]ActionResponse, len(payload))

	for i, element := range payload {
		resp, err := executeWithRetry(ctx, []ActionPayloadElement{element}, userID, false, false)
		if err != nil {
			return nil, fmt.Errorf("action %s: %w", element.Action, err)
		}
		if len(resp) > 0 {
			responses[i] = resp[0]
		}
	}

	return responses, nil
}

// ActionPayloadElement represents one action in a request payload.
type ActionPayloadElement struct {
	Action string           `json:"action"`
	Data   []map[string]any `json:"data"`
}

func executeWithRetry(ctx context.Context, payload []ActionPayloadElement, userID int, internal, atomic bool) ([]ActionResponse, error) {
	for attempt := range maxRetry {
		responses, err := executeActions(ctx, payload, userID, internal)
		if err != nil {
			var lockedErr backenderr.DatastoreLockedError
			if errors.As(err, &lockedErr) && attempt < maxRetry-1 {
				log.Printf("Datastore locked (attempt %d/%d), retrying...", attempt+1, maxRetry)
				continue
			}
			return nil, err
		}
		return responses, nil
	}
	return nil, backenderr.DatastoreLockedError{Key: "max retries exceeded"}
}

func executeActions(ctx context.Context, payload []ActionPayloadElement, userID int, internal bool) ([]ActionResponse, error) {
	ds := datastore.NewExtended()
	rm := relation.NewManager()
	var allEvents []event.Event
	responses := make([]ActionResponse, len(payload))

	for i, element := range payload {
		action, err := Lookup(element.Action)
		if err != nil {
			return nil, backenderr.ActionError{Message: err.Error()}
		}

		if action.ActionType == ActionTypeBackendInternal && !internal {
			return nil, backenderr.ActionError{
				Message: fmt.Sprintf("action %q is not allowed to be called externally", element.Action),
			}
		}

		data := make([]Instance, len(element.Data))
		for j, d := range element.Data {
			data[j] = Instance(d)
		}

		params := &ActionParams{
			UserID:          userID,
			Internal:        internal,
			Datastore:       ds,
			RelationManager: rm,
		}

		events, results, err := Perform(ctx, action, params, data)
		if err != nil {
			return nil, err
		}

		allEvents = append(allEvents, events...)
		responses[i] = ActionResponse{Results: results}
	}

	if len(allEvents) > 0 {
		wr := event.NewWriteRequest(userID)
		for _, e := range allEvents {
			wr.AddEvent(e)
		}

		// Add history information.
		for _, e := range allEvents {
			info := history.BuildInformation(e)
			if len(info) > 0 {
				wr.AddInformation(e.FQID, info...)
			}
		}

		// TODO: Write to database via datastore writer.
		_ = wr
	}

	return responses, nil
}
