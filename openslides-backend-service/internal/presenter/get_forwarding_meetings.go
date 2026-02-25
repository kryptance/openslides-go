package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getForwardingMeetingsPresenter struct{}

type getForwardingMeetingsRequest struct {
	MeetingID   int `json:"meeting_id"`
	CommitteeID int `json:"committee_id"`
}

func (p *getForwardingMeetingsPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getForwardingMeetingsRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.MeetingID == 0 {
		return nil, fmt.Errorf("meeting_id is required")
	}

	if req.CommitteeID == 0 {
		return nil, fmt.Errorf("committee_id is required")
	}

	// TODO: Verify the committee is a valid forwarding target for the meeting.
	// TODO: Fetch all meetings within the given committee and return their info.
	return []map[string]any{}, nil
}

func init() {
	RegisterPresenter("get_forwarding_meetings", &getForwardingMeetingsPresenter{})
}
