package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-ocf/kit/security/jwt"
	netHttp "net/http"
	"strings"
)

type Claims = interface{ Valid() error }
type ClaimsFunc = func(ctx context.Context, method, uri string) Claims
type OnUnauthorizedAccessFunc = func(ctx context.Context, w netHttp.ResponseWriter, r *netHttp.Request, err error)

const bearerKey = "bearer"
const authorizationKey = "authorization"

func CtxWithToken(ctx context.Context, token string) context.Context {
	if !strings.HasPrefix(strings.ToLower(token), bearerKey+" ") {
		token = fmt.Sprintf("%s %s", bearerKey, token)
	}
	return context.WithValue(ctx, authorizationKey, token)
}

func TokenFromCtx(ctx context.Context) (string, error) {
	val := ctx.Value(authorizationKey)
	if bearer, ok := val.(string); ok && strings.HasPrefix(strings.ToLower(bearer), bearerKey+" ") {
		token := bearer[7:]
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
			return nil, fmt.Errorf("invalid token: %w", err)
		}
		return ctx, nil
	}
}

func CreateAuthMiddleware(authInterceptor Interceptor, onUnauthorizedAccessFunc OnUnauthorizedAccessFunc) func(next netHttp.Handler) netHttp.Handler {
	return func(next netHttp.Handler) netHttp.Handler {
		return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, r *netHttp.Request) {
			switch r.RequestURI {
			case "/": // health check
				next.ServeHTTP(w, r)
			default:
				ctx := CtxWithToken(r.Context(), r.Header.Get("Authorization"))
				_, err := authInterceptor(ctx, r.Method, r.RequestURI)
				if err != nil {
					onUnauthorizedAccessFunc(ctx, w, r, err)
					return
				}
				next.ServeHTTP(w, r)
			}
		})
	}
}
