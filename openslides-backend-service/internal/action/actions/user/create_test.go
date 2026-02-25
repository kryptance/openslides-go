package user

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func TestCreateSimple(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "test_Xcdfgee",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"username": "test_Xcdfgee"})
}

func TestCreateWithFirstAndLastName(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"first_name": "John",
		"last_name":  "Smith",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"first_name": "John",
		"last_name":  "Smith",
	})
}

func TestCreateWithEmail(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "testuser",
		"email":    "valid@example.com",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username": "testuser",
		"email":    "valid@example.com",
	})
}

func TestCreateBrokenEmail(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "test_Xcdfgee",
		"email":    "broken@@",
	})
	tc.AssertError(resp, err)
	tc.AssertErrorContains(err, "email")
}

func TestCreateWithGender(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username":  "testuser",
		"gender_id": 1,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username":  "testuser",
		"gender_id": "1",
	})
}

func TestCreateWithDefaultPassword(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username":         "testuser",
		"default_password": "password123",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username":         "testuser",
		"default_password": "password123",
	})
}

func TestCreateWithTitle(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "testuser",
		"title":    "Dr.",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username": "testuser",
		"title":    "Dr.",
	})
}

func TestCreateWithPronoun(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "testuser",
		"pronoun":  "Test",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username": "testuser",
		"pronoun":  "Test",
	})
}

func TestCreateWithIsActive(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username":  "testuser",
		"is_active": true,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username":  "testuser",
		"is_active": "true",
	})
}

func TestCreateWithOML(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username":                      "testuser",
		"organization_management_level": "can_manage_users",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username":                      "testuser",
		"organization_management_level": "can_manage_users",
	})
}

func TestCreateMultipleUsers(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username": "user1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{"username": "user1"})

	resp2, err2 := tc.Request("user.create", map[string]any{
		"username": "user2",
	})
	tc.AssertSuccess(resp2, err2)
	tc.AssertModelExists("user/2", map[string]any{"username": "user2"})
}

func TestCreateWithAllFields(t *testing.T) {
	tc := testutil.New(t)
	tc.SetModels(map[string]map[string]any{
		"organization/1": {"name": "Test Organization"},
	})
	resp, err := tc.Request("user.create", map[string]any{
		"username":                      "testuser",
		"first_name":                    "First",
		"last_name":                     "Last",
		"email":                         "test@example.com",
		"title":                         "Prof.",
		"pronoun":                       "they",
		"gender_id":                     2,
		"default_password":              "secret",
		"is_active":                     true,
		"organization_management_level": "can_manage_users",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("user/1", map[string]any{
		"username":                      "testuser",
		"first_name":                    "First",
		"last_name":                     "Last",
		"email":                         "test@example.com",
		"title":                         "Prof.",
		"pronoun":                       "they",
		"organization_management_level": "can_manage_users",
	})
}
