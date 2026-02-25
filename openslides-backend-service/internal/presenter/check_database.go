package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type checkDatabasePresenter struct{}

type checkDatabaseRequest struct {
	Collection string `json:"collection"`
}

func (p *checkDatabasePresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req checkDatabaseRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.Collection == "" {
		return nil, fmt.Errorf("collection is required")
	}

	// TODO: Validate collection data consistency against model definition.
	// TODO: Check required fields, relation integrity, and type constraints.
	return map[string]any{"ok": true, "errors": []string{}}, nil
}

func init() {
	RegisterPresenter("check_database", &checkDatabasePresenter{})
}
