package routes

import (
	c "go-training-restful/controller"
	// m "go-training-restful/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// m.LogMIddleware(e)

	e.GET("/users", c.GetUsersController)

	return e
}
