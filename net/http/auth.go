package http

import (
	"strings"
	"fmt"
	"crypto/tls"
	"context"
	"github.com/go-ocf/kit/security/jwt"
)


type Claims = interface{ Valid() error }
type ClaimsFunc = func(ctx context.Context, method, uri string) Claims

const bearerKey = "bearer"
const authorizationKey = "authorization"

func CtxWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authorizationKey, fmt.Sprintf("%s %s", bearerKey, token))
}

func TokenFromCtx(ctx context.Context) (string, error) {
	val := ctx.Value(authorizationKey)
	if bearer, ok := val.(string); ok && strings.HasPrefix(bearer, bearerKey + " ") {
		token := strings.TrimPrefix(bearer, bearerKey+" ")
		if token == "" {
			return "", fmt.Errorf("invalid token")
		}
		return token, nil
	}
	return "", fmt.Errorf("token not found")
}

func ValidateJWT(jwksUrl string, tls tls.Config, claims ClaimsFunc) Interceptor {
	validator := jwt.NewValidator(jwksUrl, tls)
	return func(ctx context.Context, method, uri string) (context.Context, error) {
		token, err := TokenFromCtx(ctx)
		if err != nil {
			return nil, err
		}
		err = validator.ParseWithClaims(token, claims(ctx, method, uri))
		if err != nil {
			return nil, fmt.Errorf("invalid token: %v", err)
		}
		return ctx, nil
	}
}