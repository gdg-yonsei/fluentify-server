package middleware

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if err := checkAuthHeader(authHeader); err != nil {
			return c.String(err.(*echo.HTTPError).Code, err.Error())
		}
		idToken := authHeader[len("Bearer "):]

		client := getAuthClient()
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("uid", token.UID)

		return next(c)
	}
}

func checkAuthHeader(authHeader string) error {
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header needed")
	}
	if len(authHeader) <= len("Bearer ") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
	return nil
}

// TODO: DI 로 빼기
func getAuthClient() *auth.Client {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
	}

	return client
}
