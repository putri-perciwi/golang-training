package controllers

import (
	"net/http"

	"go-training-restful/lib/database"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
