package mixin

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
)

// WithTextHash wraps the UpdateInstance hook to compute a hash of the text
// field and store it as text_hash.
//
// This replaces Python's TextHashMixin.
func WithTextHash(a *action.BaseAction, textField, hashField string) {
	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		if text, ok := instance[textField]; ok {
			if textStr, ok := text.(string); ok && textStr != "" {
				hash := sha256.Sum256([]byte(textStr))
				instance[hashField] = fmt.Sprintf("%x", hash)
			} else {
				instance[hashField] = nil
			}
		}

		return origUpdate(ctx, params, instance)
	}
}
