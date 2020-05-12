package coap

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ocf/go-coap/v2/message/codes"
	"github.com/go-ocf/kit/security/jwt"
)

type Claims = interface{ Valid() error }
type ClaimsFunc = func(ctx context.Context, code codes.Code, path string) Claims

const bearerKey = "bearer"
const authorizationKey = "authorization"

func CtxWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authorizationKey, fmt.Sprintf("%s %s", bearerKey, token))
}

func TokenFromCtx(ctx context.Context) (string, error) {
	val := ctx.Value(authorizationKey)
	if bearer, ok := val.(string); ok && strings.HasPrefix(bearer, bearerKey+" ") {
		token := strings.TrimPrefix(bearer, bearerKey+" ")
		if token == "" {
			return "", fmt.Errorf("invalid token")
		}
		return token, nil
	}
	return "", fmt.Errorf("token not found")
}

func ValidateJWT(jwksURL string, tls *tls.Config, claims ClaimsFunc) Interceptor {
	validator := jwt.NewValidator(jwksURL, tls)
	return func(ctx context.Context, code codes.Code, path string) (context.Context, error) {
		token, err := TokenFromCtx(ctx)
		if err != nil {
			return nil, err
		}
		err = validator.ParseWithClaims(token, claims(ctx, code, path))
		if err != nil {
			return nil, fmt.Errorf("invalid token: %w", err)
		}
		return ctx, nil
	}
}
