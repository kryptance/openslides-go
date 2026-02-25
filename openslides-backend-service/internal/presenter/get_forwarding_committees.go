package presenter

import (
	"context"
	"encoding/json"
	"fmt"
)

type getForwardingCommitteesPresenter struct{}

type getForwardingCommitteesRequest struct {
	MeetingID int `json:"meeting_id"`
}

func (p *getForwardingCommitteesPresenter) Handle(ctx context.Context, userID int, data json.RawMessage) (any, error) {
	var req getForwardingCommitteesRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}

	if req.MeetingID == 0 {
		return nil, fmt.Errorf("meeting_id is required")
	}

	// TODO: Check that the requesting user has perm.MotionCanForward in the given meeting.
	// TODO: Resolve the meeting's committee, then find all committees that are
	//       configured as forwarding targets and return their info.
	return []map[string]any{}, nil
}

func init() {
	RegisterPresenter("get_forwarding_committees", &getForwardingCommitteesPresenter{})
}
