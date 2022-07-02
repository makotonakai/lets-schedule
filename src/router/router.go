package router

import (
	// "net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"github.com/MakotoNakai/lets-schedule/handlers"
	"github.com/MakotoNakai/lets-schedule/controllers"
)

func Initialize() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")
	// versionを取得して埋め込みして、version非依存にする
	api.GET("/", handlers.Accessible)
	api.POST("/login", handlers.Login)

	r := e.Group("/api/restricted")
	config := middleware.JWTConfig{
		Claims:     &handlers.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handlers.Restricted)

	r.GET("/users", controllers.GetUsers)
	r.POST("/users/new", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/meetings/:user_id", controllers.GetMeetingsByUserId)
	r.POST("/meetings/new", controllers.CreateMeeting)
	r.GET("/meetings/host/confirmed/:user_id", controllers.GetConfirmedMeetingsForHost)
	r.GET("/meetings/host/not-yet-confirmed/:user_id", controllers.GetNotYetConfirmedMeetingsForHost)
	r.GET("/meetings/guest/confirmed/:user_id", controllers.GetConfirmedMeetingsForGuest)
	r.GET("/meetings/guest/responded/:user_id", controllers.GetRespondedMeetingsForGuest)
	r.GET("/meetings/guest/not-yet-responded/:user_id", controllers.GetNotYetRespondedMeetingsForGuest)
	r.PUT("/meetings/:id", controllers.UpdateMeeting)
	r.DELETE("/meetings/:id", controllers.DeleteMeeting)

	r.POST("/participants/new", controllers.CreateParticipant)
	r.GET("/participants/:id", controllers.GetParticipant)
	r.PUT("/participants/:id", controllers.UpdateParticipant)
	r.DELETE("/participants/:id", controllers.DeleteParticipant)

	r.GET("/candidate_times/:user_id/:meeting_id", controllers.GetCandidateTimeByUserIdAndMeetingId)
	r.POST("/candidate_times/new", controllers.CreateCandidateTimeList)
	r.PUT("/candidate_times/:id", controllers.UpdateCandidateTime)
	r.DELETE("/candidate_times/:id", controllers.DeleteCandidateTime)

	return e

}