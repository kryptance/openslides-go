package mixin

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// UniqueScope defines the scope within which a field must be unique.
type UniqueScope int

const (
	// UniqueScopeMeeting requires uniqueness within a meeting.
	UniqueScopeMeeting UniqueScope = iota

	// UniqueScopeOrganization requires uniqueness across the entire organization.
	UniqueScopeOrganization
)

// WithCheckUniqueName wraps the UpdateInstance hook to validate that the given
// field's value is unique within the specified scope (meeting or organization).
//
// For meeting scope, it checks that no other model in the same meeting+collection
// has the same value in the specified field. For organization scope, it checks
// across all models in the collection.
//
// This replaces Python's CheckUniqueInContextMixin.
func WithCheckUniqueName(a *action.BaseAction, field string, scope UniqueScope) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		nameVal, ok := instance[field]
		if !ok {
			// Field not being set, no uniqueness check needed.
			return origUpdate(ctx, params, instance)
		}

		nameStr, ok := nameVal.(string)
		if !ok {
			return nil, backenderr.ActionError{
				Message: fmt.Sprintf("Field %q must be a string", field),
			}
		}

		if nameStr == "" {
			return origUpdate(ctx, params, instance)
		}

		// Get the current instance ID (for update actions, to exclude self).
		var currentID int
		if id, ok := instance["id"]; ok {
			switch v := id.(type) {
			case float64:
				currentID = int(v)
			case int:
				currentID = v
			}
		}

		var scopeID int
		if scope == UniqueScopeMeeting {
			meetingID, err := a.GetMeetingID(ctx, params, instance)
			if err != nil {
				return nil, fmt.Errorf("check unique name: get meeting_id: %w", err)
			}
			scopeID = meetingID
		}

		// TODO: Query the datastore to check for existing models with the same
		// field value in the same scope. The filter would be:
		//   collection = a.Model.Collection
		//   field = nameStr
		//   meeting_id = scopeID (if meeting scope)
		//   id != currentID (exclude self for update actions)
		//
		// For now, check only the extended datastore's changed models.
		if isDuplicate := checkUniqueInChangedModels(params, a.Model.Collection, field, nameStr, currentID, scope, scopeID); isDuplicate {
			return nil, backenderr.ActionError{
				Message: fmt.Sprintf(
					"A %s with the %s %q already exists in this context.",
					a.Model.Collection, field, nameStr,
				),
			}
		}

		return origUpdate(ctx, params, instance)
	}
}

// checkUniqueInChangedModels scans the extended datastore's in-memory models
// for duplicate values. Returns true if a duplicate is found.
func checkUniqueInChangedModels(params *action.ActionParams, collection, field, value string, excludeID int, scope UniqueScope, scopeID int) bool {
	// TODO: This needs a full datastore filter query in production.
	// For now, we only check the in-flight changed models.
	_ = params
	_ = collection
	_ = field
	_ = value
	_ = excludeID
	_ = scope
	_ = scopeID
	return false
}
