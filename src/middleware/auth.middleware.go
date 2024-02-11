package middleware

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthMiddleware struct {
	authClient *auth.Client
}

func (m *AuthMiddleware) Verify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if err := m.checkAuthHeader(authHeader); err != nil {
				return c.String(err.(*echo.HTTPError).Code, err.Error())
			}
			idToken := authHeader[len("Bearer "):]

			token, err := m.authClient.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return c.String(http.StatusUnauthorized, "Invalid token")
			}

			c.Set("uid", token.UID)

			return next(c)
		}
	}
}

func (m *AuthMiddleware) checkAuthHeader(authHeader string) error {
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header needed")
	}
	if len(authHeader) <= len("Bearer ") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	return nil
}

func NewAuthMiddleware(authClient *auth.Client) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}
