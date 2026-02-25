// Package permission provides a bridge to the perm package for action permission checks.
package permission

import (
	"context"
	"fmt"

	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
	"github.com/OpenSlides/openslides-go/perm"
)

// Superadmin user IDs have all permissions. Internal calls (userID == -1) bypass.
const internalUserID = -1

// HasPerm checks if the user in the context has the given permission for the meeting.
func HasPerm(ctx context.Context, userID, meetingID int, permission perm.TPermission) error {
	if userID == internalUserID {
		return nil
	}

	// TODO: Create dsfetch.Fetch from datastore and call perm.Has().
	// Pseudocode:
	//   fetch := dsfetch.New(db)
	//   allowed, err := perm.Has(ctx, fetch, userID, meetingID, permission)
	//   if err != nil { return err }
	//   if !allowed { return backenderr.MissingPermission{...} }
	_ = userID
	_ = meetingID
	_ = permission
	return nil
}

// HasOML checks if the user has the required organization management level.
func HasOML(ctx context.Context, userID int, level perm.OrganizationManagementLevel) error {
	if userID == internalUserID {
		return nil
	}

	// TODO: Fetch user's OML from datastore and compare.
	// Pseudocode:
	//   fetch := dsfetch.New(db)
	//   userOML, err := fetch.User_OrganizationManagementLevel(userID).Value(ctx)
	//   if err != nil { return err }
	//   if !isOMLSufficient(userOML, level) { return backenderr.MissingPermission{...} }
	_ = userID
	_ = level
	return nil
}

// IsSuperadmin checks if the user is a superadmin.
func IsSuperadmin(ctx context.Context, userID int) (bool, error) {
	if userID == internalUserID {
		return true, nil
	}

	// TODO: Fetch from datastore.
	// Pseudocode:
	//   fetch := dsfetch.New(db)
	//   oml, err := fetch.User_OrganizationManagementLevel(userID).Value(ctx)
	//   return oml == "superadmin", err
	return false, nil
}

// AssertNotAnonymous returns an error if the user is anonymous (userID == 0).
func AssertNotAnonymous(userID int) error {
	if userID == 0 {
		return backenderr.AnonymousNotAllowed{}
	}
	return nil
}

// CheckPermission checks a generic permission value which can be:
// - perm.TPermission: a meeting-level permission
// - perm.OrganizationManagementLevel: an organization-level permission
// - string "superadmin": requires superadmin OML
// - nil: no permission required
func CheckPermission(ctx context.Context, userID, meetingID int, permission any) error {
	if permission == nil {
		return nil
	}

	if userID == internalUserID {
		return nil
	}

	switch p := permission.(type) {
	case perm.TPermission:
		return HasPerm(ctx, userID, meetingID, p)

	case perm.OrganizationManagementLevel:
		return HasOML(ctx, userID, p)

	case string:
		if p == "superadmin" {
			isSuperadmin, err := IsSuperadmin(ctx, userID)
			if err != nil {
				return err
			}
			if !isSuperadmin {
				return backenderr.MissingPermission{
					Permission: "superadmin",
				}
			}
			return nil
		}
		return fmt.Errorf("unknown permission string %q", p)

	default:
		return fmt.Errorf("unknown permission type %T", permission)
	}
}
