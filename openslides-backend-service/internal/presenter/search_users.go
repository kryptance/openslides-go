package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type searchUsersPresenter struct{}

type searchUsersCriteria struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type searchUsersRequest struct {
	PermissionType string                `json:"permission_type"`
	PermissionID   int                   `json:"permission_id"`
	Search         []searchUsersCriteria `json:"search"`
}

func (p *searchUsersPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req searchUsersRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if len(req.Search) == 0 {
		return nil, fmt.Errorf("search criteria must not be empty")
	}

	// TODO: For each search entry, find users matching any combination of the
	//       provided fields (username, first_name, last_name, email).
	// TODO: Filter results by the requesting user's permissions scoped by
	//       permission_type and permission_id.
	results := make([][]map[string]any, len(req.Search))
	for i := range req.Search {
		results[i] = []map[string]any{}
	}

	return results, nil
}

func init() {
	RegisterPresenter("search_users", &searchUsersPresenter{})
}
