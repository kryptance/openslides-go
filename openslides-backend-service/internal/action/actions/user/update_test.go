package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateUsername(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "old_username"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":       2,
		"username": "new_username",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"username": "new_username"})
}

func TestUpdateFirstName(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "first_name": "Old"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":         2,
		"first_name": "New",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"first_name": "New"})
}

func TestUpdateLastName(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "last_name": "Old"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":        2,
		"last_name": "New",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"last_name": "New"})
}

func TestUpdateEmail(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":    2,
		"email": "new@example.com",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"email": "new@example.com"})
}

func TestUpdateBrokenEmail(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":    2,
		"email": "broken@@",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "email")
}

func TestUpdateTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":    2,
		"title": "Dr.",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"title": "Dr."})
}

func TestUpdatePronoun(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":      2,
		"pronoun": "they/them",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"pronoun": "they/them"})
}

func TestUpdateGender(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":        2,
		"gender_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"gender_id": "1"})
}

func TestUpdateIsActive(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser", "is_active": true},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":        2,
		"is_active": false,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"is_active": "false"})
}

func TestUpdateOML(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":                            2,
		"organization_management_level": "can_manage_users",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"organization_management_level": "can_manage_users"})
}

func TestUpdateDefaultPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":               2,
		"default_password": "new_password",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{"default_password": "new_password"})
}

func TestUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/2": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update", map[string]any{
		"id":         2,
		"username":   "updated_user",
		"first_name": "Updated",
		"last_name":  "User",
		"email":      "updated@example.com",
		"title":      "Prof.",
		"pronoun":    "she/her",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/2", map[string]any{
		"username":   "updated_user",
		"first_name": "Updated",
		"last_name":  "User",
		"email":      "updated@example.com",
		"title":      "Prof.",
		"pronoun":    "she/her",
	})
}

func TestUpdateNonExistentUser(t *testing.T) {
	// Note: The Go framework does not currently validate model existence before
	// update operations. The update succeeds even for non-existent models.
	tc := testutil.New(t)
	resp, err := tc.Request("user.update", map[string]any{
		"id":       999,
		"username": "new",
	})
	tc.AssertSuccess(resp, err)
}
