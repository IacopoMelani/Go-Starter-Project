package controllers

import (
	durationdata "github.com/Go-Starter-Project/models/duration_data"
	"github.com/Go-Starter-Project/models/table_record/table"

	"github.com/labstack/echo"
)

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
