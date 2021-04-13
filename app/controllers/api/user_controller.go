package api

import (
	"net/http"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/app/controllers"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/app/models/duration_data"
	"github.com/IacopoMelani/Go-Starter-Project/app/models/table"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// JwtCustomClaims - Define custom claims for token generation
type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// GetAllUser - Return all users
func GetAllUser(c echo.Context) error {

	userList, err := table.LoadAllUsers(db.GetConnection())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controllers.FailedResponse(c, 1, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, controllers.SuccessResponse(c, echo.Map{
		"users": userList,
	}))
}

// GetDurataionUsers - Return users retrived with pkg/model/duration_data
func GetDurataionUsers(c echo.Context) error {
	data := durationdata.GetUsersData()
	return c.JSON(http.StatusOK, controllers.SuccessResponse(c, echo.Map{
		"data": data.GetSafeContent(),
	}))
}

// Login - Define login controller
func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "Mario" || password != "123456" {
		return c.JSON(http.StatusUnauthorized, controllers.FailedResponse(c, 1, "Wrong credentials", nil))
	}

	claims := &JwtCustomClaims{
		"Mario",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	generatedToken, err := token.SignedString([]byte("secret here")) // maybe reads from .env file
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controllers.FailedResponse(c, 2, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, controllers.SuccessResponse(c, echo.Map{
		"token": generatedToken,
	}))
}
