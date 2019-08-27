package controllers

import (
	durationdata "Go-Starter-Project/models/duration_data"
	"Go-Starter-Project/models/table_record/table"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

// JwtCustomClaims - Definisce il custom claims per la generazione del token
type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// GetAllUser - Restituisce tutti gli utenti
func GetAllUser(c echo.Context) error {

	userList, err := table.LoadAllUsers()
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

// GetDurataionUsers - Restituisce gli utenti recuperati con DurationData
func GetDurataionUsers(c echo.Context) error {

	data := durationdata.GetUsersData()

	return c.JSON(200, Response{
		Status:  0,
		Success: true,
		Message: "ok!",
		Content: data.Content,
	})
}

// Login - Si occupa di effettuare il login
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
