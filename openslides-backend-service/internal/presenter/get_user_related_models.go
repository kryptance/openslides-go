package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getUserRelatedModelsPresenter struct{}

type getUserRelatedModelsRequest struct {
	UserIDs []int `json:"user_ids"`
}

func (p *getUserRelatedModelsPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getUserRelatedModelsRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if len(req.UserIDs) == 0 {
		return nil, fmt.Errorf("user_ids is required and must not be empty")
	}

	// TODO: For each user, collect all related models (meetings, committees,
	//       assigned motions, speaker entries, assignments, etc.).
	result := make(map[int]map[string]any, len(req.UserIDs))
	for _, uid := range req.UserIDs {
		result[uid] = map[string]any{
			"meetings":   []map[string]any{},
			"committees": []map[string]any{},
		}
	}

	return result, nil
}

func init() {
	RegisterPresenter("get_user_related_models", &getUserRelatedModelsPresenter{})
}
