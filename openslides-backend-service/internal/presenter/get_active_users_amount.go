package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getActiveUsersAmountPresenter struct{}

type getActiveUsersAmountRequest struct{}

func (p *getActiveUsersAmountPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getActiveUsersAmountRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if userID == 0 {
		return nil, fmt.Errorf("anonymous users are not allowed to use this presenter")
	}

	// TODO: Query the datastore for all users where is_active is true.
	// TODO: Return the count of active users.
	return map[string]any{"active_users_amount": 0}, nil
}

func init() {
	RegisterPresenter("get_active_users_amount", &getActiveUsersAmountPresenter{})
}
