package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/MakotoNakai/lets-schedule/controllers"
)

func Router() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// versionを取得して埋め込みして、version非依存にする
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping")
	})

	e.GET("/users", controllers.GetUsers)
	e.POST("/users/new", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	return e

}