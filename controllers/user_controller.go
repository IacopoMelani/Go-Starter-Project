package controllers

import (
	"Go-Starter-Project/models"

	"github.com/labstack/echo"
)

// GetAllUser - Restituisce tutti gli utenti
func GetAllUser(c echo.Context) error {

	userList, err := models.LoadAllUser()
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
