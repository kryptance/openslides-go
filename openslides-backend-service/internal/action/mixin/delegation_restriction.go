package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// DelegationType defines which delegation restriction to enforce.
type DelegationType int

const (
	// DelegationVote restricts voting when the user has delegated their vote.
	DelegationVote DelegationType = iota

	// DelegationSpeaker restricts speaking when the user has delegated their vote.
	DelegationSpeaker
)

// WithDelegationRestriction wraps the CheckPermissions hook to enforce
// delegation-based restrictions. When a user has delegated their vote to
// another user in a meeting, certain operations (voting, speaking) may be
// forbidden depending on the meeting's delegation settings.
//
// The meeting settings checked are:
//   - users_forbid_delegator_as_voter (for DelegationVote)
//   - users_forbid_delegator_as_speaker (for DelegationSpeaker)
//
// This replaces Python's VoteDelegationMixin / check_for_delegation logic.
func WithDelegationRestriction(a *action.BaseAction, delegationType DelegationType) {
	origCheck := a.CheckPermissions
	a.CheckPermissions = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		// First, run the original permission check.
		if err := origCheck(ctx, params, instance); err != nil {
			return err
		}

		// Determine the meeting.
		meetingID, err := a.GetMeetingID(ctx, params, instance)
		if err != nil {
			return fmt.Errorf("delegation restriction: get meeting_id: %w", err)
		}
		if meetingID == 0 {
			return nil
		}

		// Determine which setting to check based on delegation type.
		var settingField string
		switch delegationType {
		case DelegationVote:
			settingField = "users_forbid_delegator_as_voter"
		case DelegationSpeaker:
			settingField = "users_forbid_delegator_as_speaker"
		}

		// Fetch the meeting setting.
		meeting, err := params.Datastore.Get("meeting", meetingID, []string{settingField})
		if err != nil {
			return fmt.Errorf("delegation restriction: fetch meeting %d: %w", meetingID, err)
		}

		forbid, ok := meeting[settingField]
		if !ok || forbid != true {
			// Setting is not enabled, no restriction.
			return nil
		}

		// Check if the acting user has delegated their vote in this meeting.
		// We need to find the meeting_user record for the acting user.
		// TODO: Query the datastore:
		//   meeting_user WHERE user_id = params.UserID AND meeting_id = meetingID
		//   Check if vote_delegated_to_id is set.
		hasDelegation, err := checkUserHasDelegation(params, meetingID)
		if err != nil {
			return fmt.Errorf("delegation restriction: check delegation: %w", err)
		}

		if hasDelegation {
			var actionDesc string
			switch delegationType {
			case DelegationVote:
				actionDesc = "vote"
			case DelegationSpeaker:
				actionDesc = "speak"
			}
			return backenderr.ActionError{
				Message: fmt.Sprintf(
					"You have delegated your vote in meeting %d and are therefore not allowed to %s.",
					meetingID, actionDesc,
				),
			}
		}

		return nil
	}
}

// checkUserHasDelegation checks whether the acting user has delegated their vote
// in the specified meeting.
func checkUserHasDelegation(params *action.ActionParams, meetingID int) (bool, error) {
	// TODO: Look up the meeting_user for this user and meeting, then check
	// if vote_delegated_to_id is set and non-nil.
	_ = params
	_ = meetingID
	return false, nil
}
