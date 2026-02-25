package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getMediafileContextPresenter struct{}

type getMediafileContextRequest struct {
	MediafileID int `json:"mediafile_id"`
}

func (p *getMediafileContextPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getMediafileContextRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.MediafileID == 0 {
		return nil, fmt.Errorf("mediafile_id is required")
	}

	// TODO: Look up the mediafile by ID and determine where it is used.
	// TODO: Check relations like logo, font, attachment on topics/motions, etc.
	return map[string]any{
		"mediafile_id": req.MediafileID,
		"used_in":      []map[string]any{},
	}, nil
}

func init() {
	RegisterPresenter("get_mediafile_context", &getMediafileContextPresenter{})
}
