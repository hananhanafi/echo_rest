package routes

import (
	"echo-rest/controllers"
	"echo-rest/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/employee", controllers.FetchAllEmployee, middleware.IsAuthenticated)
	e.POST("/employee", controllers.StoreEmployee, middleware.IsAuthenticated)
	e.PUT("/employee", controllers.UpdateEmployee, middleware.IsAuthenticated)
	e.DELETE("/employee", controllers.DeleteEmployee, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.POST("/login/custom", controllers.CheckLoginCustom)

	return e
}
