package keycloak

import (
	"context"
	"encoding/json"

	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/vanti-dev/sc-bos/pkg/auth/jwks"
	"github.com/vanti-dev/sc-bos/pkg/auth/token"

	"github.com/vanti-dev/sc-bos/pkg/auth"
)

type accessTokenPayload struct {
	jwt.Claims
	Roles          []string                  `json:"roles"`
	Scopes         auth.JWTScopes            `json:"scope"`
	ResourceAccess map[string]resourceAccess `json:"resource_access"`
}

func (a *accessTokenPayload) AllRoles() []string {
	var roles []string
	roles = append(roles, a.Roles...)
	for _, res := range a.ResourceAccess {
		roles = append(roles, res.Roles...)
	}
	return roles
}

func (a *accessTokenPayload) IsAppOnly() bool {
	return false
}

type resourceAccess struct {
	Roles []string `json:"roles"`
}

func NewTokenVerifier(config *Config, keySet jwks.KeySet) *TokenValidator {
	return &TokenValidator{
		keySet: keySet,
		expected: jwt.Expected{
			Audience: jwt.Audience{config.ClientID},
			Issuer:   config.Issuer(),
		},
	}
}

type TokenValidator struct {
	keySet   jwks.KeySet
	expected jwt.Expected
}

func (v *TokenValidator) ValidateAccessToken(ctx context.Context, tokenStr string) (*token.Claims, error) {
	payloadBytes, err := v.keySet.VerifySignature(ctx, tokenStr)
	if err != nil {
		return nil, err
	}

	var payload accessTokenPayload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return nil, err
	}

	err = payload.Claims.Validate(v.expected)
	if err != nil {
		return nil, err
	}

	return &token.Claims{
		Roles:     payload.AllRoles(),
		Scopes:    payload.Scopes,
		IsService: payload.IsAppOnly(),
	}, nil
}
