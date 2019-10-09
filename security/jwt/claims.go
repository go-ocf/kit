package jwt

import (
	"fmt"
	"time"

	"github.com/go-ocf/kit/strings"
)

var TimeFunc = time.Now

type Claims struct {
	ClientID string   `json:"client_id"`
	Email    string   `json:"email"`
	Scope    []string `json:"scope"`
	StandardClaims
}

// https://tools.ietf.org/html/rfc7519#section-4.1
type StandardClaims struct {
	Audience  interface{} `json:"aud,omitempty"`
	ExpiresAt int64       `json:"exp,omitempty"`
	Id        string      `json:"jti,omitempty"`
	IssuedAt  int64       `json:"iat,omitempty"`
	Issuer    string      `json:"iss,omitempty"`
	NotBefore int64       `json:"nbf,omitempty"`
	Subject   string      `json:"sub,omitempty"`
}

func (c StandardClaims) GetAudience() []string {
	return strings.ToSlice(c.Audience)
}

func (c StandardClaims) Valid() error {
	now := TimeFunc().Unix()
	if now > c.ExpiresAt {
		return fmt.Errorf("token is expired")
	}
	if now < c.IssuedAt {
		return fmt.Errorf("token used before issued")
	}
	if now < c.NotBefore {
		return fmt.Errorf("token is not valid yet")
	}
	return nil
}
