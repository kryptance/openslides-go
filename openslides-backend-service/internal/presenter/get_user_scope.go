package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getUserScopePresenter struct{}

type getUserScopeRequest struct {
	UserIDs []int `json:"user_ids"`
}

func (p *getUserScopePresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getUserScopeRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if len(req.UserIDs) == 0 {
		return nil, fmt.Errorf("user_ids is required and must not be empty")
	}

	// TODO: For each user, determine the permission scope of the requesting user.
	// TODO: Scope is one of "organization", "committee", or "meeting" depending
	//       on where the requesting user has management permissions over the target user.
	result := make(map[int]map[string]any, len(req.UserIDs))
	for _, uid := range req.UserIDs {
		result[uid] = map[string]any{
			"collection":    "organization",
			"id":            1,
			"user_can_update": false,
		}
	}

	return result, nil
}

func init() {
	RegisterPresenter("get_user_scope", &getUserScopePresenter{})
}
