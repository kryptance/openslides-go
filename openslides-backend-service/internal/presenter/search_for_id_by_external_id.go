package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type searchForIDByExternalIDPresenter struct{}

type searchForIDByExternalIDRequest struct {
	Collection string `json:"collection"`
	ExternalID string `json:"external_id"`
	ContextID  int    `json:"context_id"`
}

func (p *searchForIDByExternalIDPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req searchForIDByExternalIDRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.Collection == "" {
		return nil, fmt.Errorf("collection is required")
	}

	if req.ExternalID == "" {
		return nil, fmt.Errorf("external_id is required")
	}

	// TODO: Search the given collection for a model whose external_id matches,
	//       scoped by context_id if provided.
	// TODO: Return the internal ID if found, or an error if not found.
	return map[string]any{"id": 0}, nil
}

func init() {
	RegisterPresenter("search_for_id_by_external_id", &searchForIDByExternalIDPresenter{})
}
