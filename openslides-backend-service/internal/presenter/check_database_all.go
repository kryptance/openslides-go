package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type checkDatabaseAllPresenter struct{}

type checkDatabaseAllRequest struct{}

func (p *checkDatabaseAllPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req checkDatabaseAllRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	// TODO: Iterate over all known collections and check each for consistency.
	// TODO: Aggregate errors per collection into the result map.
	return map[string]any{"ok": true, "results": map[string][]string{}}, nil
}

func init() {
	RegisterPresenter("check_database_all", &checkDatabaseAllPresenter{})
}
