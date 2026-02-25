package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestUpdateSelfCorrect(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "username_srtgb123"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"username": "username_Xcdfgee",
		"email":    "email1@example.com",
		"pronoun":  "Test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username": "username_Xcdfgee",
		"email":    "email1@example.com",
		"pronoun":  "Test",
	})
}

func TestUpdateSelfUsernameOnly(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "old_name"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"username": "new_name",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"username": "new_name"})
}

func TestUpdateSelfEmailOnly(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"email": "newemail@example.com",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"email": "newemail@example.com"})
}

func TestUpdateSelfBrokenEmail(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "username_srtgb123"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"email": "broken@@",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "email")
}

func TestUpdateSelfPronoun(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"pronoun": "they/them",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"pronoun": "they/them"})
}

func TestUpdateSelfGender(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"user/1": {"username": "testuser"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"gender_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"gender_id": "1"})
}

func TestUpdateSelfSetsOwnId(t *testing.T) {
	tc := testutil.New(t)
	tc.UserID = 5
	tc.SetModels(map[string]map[string]any{
		"user/5": {"username": "user5"},
	})
	resp, err := tc.Request("user.update_self", map[string]any{
		"username": "updated_user5",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/5", map[string]any{"username": "updated_user5"})
}

func TestUpdateSelfAnonymous(t *testing.T) {
	tc := testutil.New(t)
	tc.UserID = 0 // anonymous
	resp, err := tc.Request("user.update_self", map[string]any{
		"email": "user@openslides.org",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "nonymous")
}
