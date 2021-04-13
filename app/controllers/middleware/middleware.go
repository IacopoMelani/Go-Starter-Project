package middleware

import (
	"github.com/IacopoMelani/Go-Starter-Project/app/controllers"
	"github.com/labstack/echo/v4"
)

const (
	secretToken = "THIS_IS_A_SECRET!" // Maybe loaded by env file or dynamic
)

// APIBasicAuthMiddleware - Defines a middleware for a basic auth-token based authentication
func APIBasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authToken := c.Request().Header.Get("authtoken")

		if authToken == "" || authToken != secretToken {
			return controllers.APIAuthBasicFailedResponse(c)
		}

		return next(c)
	}
}
