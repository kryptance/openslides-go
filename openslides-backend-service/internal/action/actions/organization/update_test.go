package organization

import (
	"testing"

	"github.com/OpenSlides/openslides-backend-service/internal/action/testutil"
)

func setupOrganization(tc *testutil.ActionTestCase) {
	tc.SetModels(map[string]map[string]any{
		"organization/1": {
			"name":        "aBuwxoYU",
			"description": "XrHbAWiF",
			"theme_id":    1,
			"theme_ids":   []any{1, 2},
		},
		"theme/1": {
			"name":                        "default",
			"organization_id":             1,
			"theme_for_organization_id":   1,
		},
		"theme/2": {
			"name":            "default2",
			"organization_id": 1,
		},
	})
}

func TestOrganizationUpdateNameAndDescription(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":          1,
		"name":        "testtest",
		"description": "blablabla",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"name":        "testtest",
		"description": "blablabla",
	})
}

func TestOrganizationUpdateLegalNotice(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":           1,
		"legal_notice": "GYjDABmD",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"legal_notice": "GYjDABmD",
	})
}

func TestOrganizationUpdatePrivacyPolicy(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":             1,
		"privacy_policy": "test1",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"privacy_policy": "test1",
	})
}

func TestOrganizationUpdateLoginText(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":         1,
		"login_text": "test2",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"login_text": "test2",
	})
}

func TestOrganizationUpdateTheme(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":       1,
		"theme_id": 2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"theme_id": 2,
	})
}

func TestOrganizationUpdateResetPasswordVerboseErrors(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":                            1,
		"reset_password_verbose_errors": false,
	})
	tc.AssertSuccess(resp, err)
}

func TestOrganizationUpdateEnableElectronicVoting(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":                       1,
		"enable_electronic_voting": true,
	})
	tc.AssertSuccess(resp, err)
}

func TestOrganizationUpdateLimitOfMeetings(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":                 1,
		"limit_of_meetings":  2,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"limit_of_meetings": 2,
	})
}

func TestOrganizationUpdateLimitOfUsers(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":             1,
		"limit_of_users": 100,
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"limit_of_users": 100,
	})
}

func TestOrganizationUpdateUrl(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":  1,
		"url": "https://openslides.example.com",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"url": "https://openslides.example.com",
	})
}

func TestOrganizationUpdateSamlEnabled(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":           1,
		"saml_enabled": true,
	})
	tc.AssertSuccess(resp, err)
}

func TestOrganizationUpdateSamlLoginButtonText(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":                     1,
		"saml_login_button_text": "Login with SSO",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"saml_login_button_text": "Login with SSO",
	})
}

func TestOrganizationUpdateSamlAttrMapping(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id": 1,
		"saml_attr_mapping": map[string]any{
			"saml_id":    "username",
			"title":      "title",
			"first_name": "firstName",
			"last_name":  "lastName",
			"email":      "email",
		},
	})
	tc.AssertSuccess(resp, err)
}

func TestOrganizationUpdateMissingId(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"name": "testtest",
	})
	tc.AssertError(resp, err)
}

func TestOrganizationUpdateMultipleFields(t *testing.T) {
	tc := testutil.New(t)
	setupOrganization(tc)
	resp, err := tc.Request("organization.update", map[string]any{
		"id":                        1,
		"name":                      "testtest",
		"description":               "blablabla",
		"legal_notice":              "GYjDABmD",
		"privacy_policy":            "test1",
		"login_text":                "test2",
		"reset_password_verbose_errors": false,
		"url":                       "https://openslides.example.com",
	})
	tc.AssertSuccess(resp, err)
	tc.AssertModelExists("organization/1", map[string]any{
		"name":           "testtest",
		"description":    "blablabla",
		"legal_notice":   "GYjDABmD",
		"privacy_policy": "test1",
		"login_text":     "test2",
		"url":            "https://openslides.example.com",
	})
}
