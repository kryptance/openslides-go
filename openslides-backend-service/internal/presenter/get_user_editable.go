package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getUserEditablePresenter struct{}

type getUserEditableRequest struct {
	UserIDs []int `json:"user_ids"`
}

func (p *getUserEditablePresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getUserEditableRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if len(req.UserIDs) == 0 {
		return nil, fmt.Errorf("user_ids is required and must not be empty")
	}

	// TODO: For each requested user, determine which fields the current user
	//       is permitted to edit based on organization, committee, and meeting roles.
	result := make(map[int]map[string]bool, len(req.UserIDs))
	for _, uid := range req.UserIDs {
		result[uid] = map[string]bool{
			"editable": false,
		}
	}

	return result, nil
}

func init() {
	RegisterPresenter("get_user_editable", &getUserEditablePresenter{})
}
