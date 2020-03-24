package routes

import (
	"github.com/IacopoMelani/Go-Starter-Project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitGetRoutes - Declare the GET routes
func InitGetRoutes(e *echo.Echo) {
	e.GET("user/all", controllers.GetAllUser)
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &controllers.JwtCustomClaims{},
		SigningKey: []byte("bomba"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user/duration", controllers.GetDurataionUsers)
	e.GET("user/duration", controllers.GetDurataionUsers)
}

// InitPostRoutes - Declare the POST routes
func InitPostRoutes(e *echo.Echo) {
	e.POST("/user/login", controllers.Login)
}
