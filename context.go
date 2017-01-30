package jwt

import (
	"golang.org/x/net/context"

	jwt "github.com/dgrijalva/jwt-go"
)

type contextKey int

const (
	jwtKey contextKey = iota + 1
)

// WithJWT creates a child context containing the given JWT.
func WithJWT(ctx context.Context, t *jwt.Token) context.Context {
	return context.WithValue(ctx, jwtKey, t)
}

// ContextJWT retrieves the JWT token from a `context` that went through our security middleware.
func ContextJWT(ctx context.Context) *jwt.Token {
	token, _ := ctx.Value(jwtKey).(*jwt.Token)
	return token
}