package routes

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestInitRoutes(t *testing.T) {

	e := echo.New()

	InitGetRoutes(e)
	InitPostRoutes(e)
}
