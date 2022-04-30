package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"github.com/MakotoNakai/lets-schedule/models"
	"github.com/MakotoNakai/lets-schedule/database"
	"github.com/MakotoNakai/lets-schedule/controllers"
)

func Initialize() *echo.Echo {

	db := database.Connect()
	db.AutoMigrate(&models.User{}, &models.Meeting{}, &models.Participant{}, &models.CandidateTime{})
	database.Seed(db)

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

	e.GET("/meetings", controllers.GetMeetings)
	e.POST("/meetings/new", controllers.CreateMeeting)
	e.GET("/meetings/:id", controllers.GetMeeting)
	e.PUT("/meetings/:id", controllers.UpdateMeeting)
	e.DELETE("/meetings/:id", controllers.DeleteMeeting)

	e.GET("/participants", controllers.GetParticipants)
	e.POST("/participants/new", controllers.CreateParticipant)
	e.GET("/participants/:meeting_id", controllers.GetParticipant)
	e.PUT("/participants/:meeting_id", controllers.UpdateParticipant)
	e.DELETE("/participants/:meeting_id", controllers.DeleteParticipant)

	e.GET("/candidate_times", controllers.GetParticipants)
	e.POST("/candidate_times/new", controllers.CreateParticipant)
	e.GET("/candidate_times/:id", controllers.GetParticipant)
	e.PUT("/candidate_times/:id", controllers.UpdateParticipant)
	e.DELETE("/candidate_times/:id", controllers.DeleteParticipant)

	return e

}