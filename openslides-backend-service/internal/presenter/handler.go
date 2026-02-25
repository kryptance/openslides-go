// Package presenter implements the presenter request handler.
package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

// PresenterRequest represents a single presenter call.
type PresenterRequest struct {
	Presenter string          `json:"presenter"`
	Data      json.RawMessage `json:"data"`
}

// Handle processes a list of presenter requests.
func Handle(ctx context.Context, userID int, payload []json.RawMessage) ([]any, error) {
	results := make([]any, len(payload))

	for i, raw := range payload {
		var req PresenterRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, fmt.Errorf("parsing presenter request %d: %w", i, err)
		}

		presenter, err := LookupPresenter(req.Presenter)
		if err != nil {
			return nil, err
		}

		result, err := presenter.Handle(ctx, userID, req.Data)
		if err != nil {
			return nil, fmt.Errorf("presenter %s: %w", req.Presenter, err)
		}

		results[i] = result
	}

	return results, nil
}
