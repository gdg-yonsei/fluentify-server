package middleware

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/gdsc-ys/fluentify-server/src/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	Verify() echo.MiddlewareFunc
}

type AuthMiddlewareImpl struct {
	authClient *auth.Client
}

func (m *AuthMiddlewareImpl) Verify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if err := m.checkAuthHeader(authHeader); err != nil {
				return err
			}
			splitToken := strings.Split(authHeader, "Bearer ")
			idToken := splitToken[1]

			token, err := m.authClient.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return model.NewCustomHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			c.Set("uid", token.UID)

			return next(c)
		}
	}
}

func (m *AuthMiddlewareImpl) checkAuthHeader(authHeader string) error {
	if authHeader == "" {
		return model.NewCustomHTTPError(http.StatusUnauthorized, "Authorization header needed")
	}
	if len(authHeader) <= len("Bearer ") {
		return model.NewCustomHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	return nil
}

func AuthMiddlewareInit(authClient *auth.Client) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		authClient: authClient,
	}
}
