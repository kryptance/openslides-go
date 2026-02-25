package mixin

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	backenderr "github.com/OpenSlides/openslides-backend-service/internal/errors"
)

// WithEmailCheck wraps the ValidateFields hook to validate that the given field
// contains a syntactically valid email address when set.
//
// This replaces Python's EmailCheckMixin.
func WithEmailCheck(a *action.BaseAction, emailField string) {
	origValidate := a.ValidateFields
	a.ValidateFields = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		if err := origValidate(ctx, params, instance); err != nil {
			return err
		}

		emailVal, ok := instance[emailField]
		if !ok {
			return nil
		}

		emailStr, ok := emailVal.(string)
		if !ok {
			return backenderr.ActionError{
				Message: fmt.Sprintf("Field %q must be a string", emailField),
			}
		}

		// Empty string is allowed (field is optional).
		if emailStr == "" {
			return nil
		}

		// Validate email format using Go's net/mail parser.
		if _, err := mail.ParseAddress(emailStr); err != nil {
			return backenderr.ActionError{
				Message: fmt.Sprintf("%q is not a valid email address", emailStr),
			}
		}

		return nil
	}
}

// WithEmailSenderCheck wraps the ValidateFields hook to validate that the
// email sender fields (users_email_sender, users_email_replyto) in meeting
// settings are valid.
//
// The sender field must be either a valid email address or a name token
// (alphanumeric only). The reply-to field, if set, must be a valid email address.
//
// This replaces Python's EmailSenderCheckMixin / email sender validation.
func WithEmailSenderCheck(a *action.BaseAction, senderField, replyToField string) {
	origValidate := a.ValidateFields
	a.ValidateFields = func(ctx context.Context, params *action.ActionParams, instance action.Instance) error {
		if err := origValidate(ctx, params, instance); err != nil {
			return err
		}

		// Validate sender field.
		if senderVal, ok := instance[senderField]; ok {
			senderStr, ok := senderVal.(string)
			if !ok {
				return backenderr.ActionError{
					Message: fmt.Sprintf("Field %q must be a string", senderField),
				}
			}

			if senderStr != "" {
				if !isValidEmailSender(senderStr) {
					return backenderr.ActionError{
						Message: fmt.Sprintf(
							"%q is not a valid email sender. Must be a valid email address or an alphanumeric name.",
							senderStr,
						),
					}
				}
			}
		}

		// Validate reply-to field.
		if replyToVal, ok := instance[replyToField]; ok {
			replyToStr, ok := replyToVal.(string)
			if !ok {
				return backenderr.ActionError{
					Message: fmt.Sprintf("Field %q must be a string", replyToField),
				}
			}

			if replyToStr != "" {
				if _, err := mail.ParseAddress(replyToStr); err != nil {
					return backenderr.ActionError{
						Message: fmt.Sprintf("%q is not a valid reply-to email address", replyToStr),
					}
				}
			}
		}

		return nil
	}
}

// isValidEmailSender checks whether a string is either a valid email address
// or a valid sender name token (alphanumeric, hyphens, dots, underscores).
func isValidEmailSender(s string) bool {
	// Try as email address first.
	if _, err := mail.ParseAddress(s); err == nil {
		return true
	}

	// Check as a name token: must be non-empty, alphanumeric with
	// allowed punctuation (hyphens, dots, underscores).
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	for _, c := range s {
		if !isAlphaNumericOrPunct(c) {
			return false
		}
	}
	return true
}

func isAlphaNumericOrPunct(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '-' || c == '.' || c == '_'
}
