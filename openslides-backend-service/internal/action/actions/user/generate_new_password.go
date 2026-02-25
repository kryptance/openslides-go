package user

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/OpenSlides/openslides-backend-service/internal/action"
	"github.com/OpenSlides/openslides-backend-service/internal/model"
	"github.com/OpenSlides/openslides-go/perm"
)

const (
	passwordLength  = 10
	passwordCharset = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
)

func init() {
	a := action.NewUpdateAction("user.generate_new_password", model.MustGet("user"))
	a.Permission = perm.UserCanManage

	a.Schema = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"id": map[string]any{"type": "integer"},
		},
		"required": []string{"id"},
	}

	origUpdate := a.UpdateInstance
	a.UpdateInstance = func(ctx context.Context, params *action.ActionParams, instance action.Instance) (action.Instance, error) {
		password, err := generateRandomPassword(passwordLength)
		if err != nil {
			return nil, err
		}

		// Set as both default_password (plaintext) and password (should be hashed).
		instance["default_password"] = password

		// TODO: Hash the password using bcrypt before storing.
		instance["password"] = password // placeholder until bcrypt is integrated

		return origUpdate(ctx, params, instance)
	}

	action.Register(a)
}

// generateRandomPassword creates a cryptographically random password of the given length
// using the defined charset (ambiguous characters excluded).
func generateRandomPassword(length int) (string, error) {
	charsetLen := big.NewInt(int64(len(passwordCharset)))
	result := make([]byte, length)
	for i := range result {
		idx, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = passwordCharset[idx.Int64()]
	}
	return string(result), nil
}
