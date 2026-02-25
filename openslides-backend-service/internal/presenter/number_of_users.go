package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type numberOfUsersPresenter struct{}

type numberOfUsersRequest struct {
	Filter map[string]any `json:"filter"`
}

func (p *numberOfUsersPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req numberOfUsersRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	// TODO: Count all users in the datastore, applying the optional filter.
	return map[string]any{"number_of_users": 0}, nil
}

func init() {
	RegisterPresenter("number_of_users", &numberOfUsersPresenter{})
}
