package routes

import (
	"github.com/labstack/echo"
	"testing"
)

func TestInitRoutes(t *testing.T) {

	e := echo.New()

	InitGetRoutes(e)
	InitPostRoutes(e)
}