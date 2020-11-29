package controllers

import (
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"

	durationdata "github.com/IacopoMelani/Go-Starter-Project/app/models/duration_data"
	"github.com/IacopoMelani/Go-Starter-Project/app/models/table"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
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
		return c.JSON(500, Response{
			Status:  1,
			Success: false,
			Message: err.Error(),
		})

	}

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: userList,
	})
}

// GetDurataionUsers - Return users retrived with pkg/model/duration_data
func GetDurataionUsers(c echo.Context) error {

	data := durationdata.GetUsersData()

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: data.GetSafeContent(),
	})
}

// Login - Define login controller
func Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "Mario" || password != "123456" {

		return c.JSON(401, Response{
			Status:  1,
			Success: false,
			Message: "Wrong credentials",
		})
	}

	claims := &JwtCustomClaims{
		"Mario",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	generatedToken, err := token.SignedString([]byte("bomba"))
	if err != nil {

		return c.JSON(500, Response{
			Status:  2,
			Success: false,
			Message: "Error",
		})
	}

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: echo.Map{
			"token": generatedToken,
		},
	})
}
