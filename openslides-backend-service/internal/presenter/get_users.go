package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getUsersPresenter struct{}

type getUsersRequest struct {
	StartIndex   int                    `json:"start_index"`
	Entries      int                    `json:"entries"`
	SortCriteria []string              `json:"sort_criteria"`
	Reverse      bool                   `json:"reverse"`
	Filter       map[string]any         `json:"filter"`
}

func (p *getUsersPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getUsersRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	// TODO: Apply filter criteria to the user collection.
	// TODO: Sort results by the given sort_criteria (optionally reversed).
	// TODO: Return the paginated slice from start_index with at most entries results.
	return map[string]any{
		"users": []map[string]any{},
		"total": 0,
	}, nil
}

func init() {
	RegisterPresenter("get_users", &getUsersPresenter{})
}
