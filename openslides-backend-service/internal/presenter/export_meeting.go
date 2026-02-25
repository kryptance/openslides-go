package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type exportMeetingPresenter struct{}

type exportMeetingRequest struct {
	MeetingID int `json:"meeting_id"`
}

func (p *exportMeetingPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req exportMeetingRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.MeetingID == 0 {
		return nil, fmt.Errorf("meeting_id is required")
	}

	// TODO: Verify that the requesting user is a meeting admin or superadmin.
	// TODO: Fetch all collections belonging to the meeting and serialize them.
	return map[string]any{
		"meeting": map[string]any{
			"id": req.MeetingID,
		},
	}, nil
}

func init() {
	RegisterPresenter("export_meeting", &exportMeetingPresenter{})
}
