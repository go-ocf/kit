package jwt

import (
	"crypto/tls"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Validator struct {
	keys *KeyCache
}

func NewValidator(jwksURL string, tls *tls.Config) *Validator {
	return &Validator{keys: NewKeyCache(jwksURL, tls)}
}

func (v *Validator) Parse(token string) (jwt.MapClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}
	t, err := jwt.Parse(token, v.keys.GetOrFetchKey)
	if t == nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}
	c := t.Claims.(jwt.MapClaims)
	if err != nil {
		return c, fmt.Errorf("could not parse token: %w", err)
	}
	return c, nil
}

func (v *Validator) ParseWithClaims(token string, claims jwt.Claims) error {
	if token == "" {
		return fmt.Errorf("missing token")
	}
	_, err := jwt.ParseWithClaims(token, claims, v.keys.GetOrFetchKey)
	if err != nil {
		return fmt.Errorf("could not parse token: %w", err)
	}
	return nil
}
