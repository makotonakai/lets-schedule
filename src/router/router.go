package router

import (
	// "net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/MakotoNakai/lets-schedule/controllers"
	"github.com/MakotoNakai/lets-schedule/handlers"
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
	api.POST("/signup", controllers.CreateUser)
	api.POST("/send-email", handlers.SendEmail)

	api.POST("/user/:id/reset-password", controllers.ResetPassword)

	r := e.Group("/api/restricted")
	config := middleware.JWTConfig{
		Claims:     &handlers.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handlers.Restricted)

	r.GET("/meetings/:id", controllers.GetMeetingById)
	r.GET("/meetings/user/:user_id", controllers.GetAllMeetings)
	r.GET("/meetings/host/confirmed/:user_id", controllers.GetConfirmedMeetingsForHost)
	r.GET("/meetings/host/not-confirmed/:user_id", controllers.GetNotConfirmedMeetingsForHost)
	r.GET("/meetings/host/not-responded/:user_id", controllers.GetNotRespondedMeetingsForHost)
	r.GET("/meetings/guest/confirmed/:user_id", controllers.GetConfirmedMeetingsForGuest)
	r.GET("/meetings/guest/not-confirmed/:user_id", controllers.GetNotConfirmedMeetingsForGuest)
	r.GET("/meetings/guest/not-responded/:user_id", controllers.GetNotRespondedMeetingsForGuest)
	r.POST("/meetings/new", controllers.CreateMeeting)
	r.PUT("/meetings/:id", controllers.UpdateMeetingById)

	r.GET("/candidate_times/user/:user_id/meeting/:meeting_id", controllers.GetCandidateTimeByUserIdAndMeetingId)
	r.GET("/candidate_times/meeting/:meeting_id", controllers.GetCandidateTimeWithUserNameByMeetingId)
	r.POST("/candidate_times/new", controllers.CreateCandidateTime)
	r.PUT("/candidate_times/user/:user_id/meeting/:meeting_id", controllers.UpdateCandidateTimeByUserIdAndMeetingId)
	r.GET("/candidate_times/available-time/:meeting_id", controllers.GetAvailableTimeByMeetingId)

	r.POST("/friends/new", controllers.CreateFriend)

	r.POST("/participants/new", controllers.CreateParticipant)
	r.GET("/participants/:meeting_id", controllers.GetParticipantByMeetingId)
	r.PUT("/participants/meeting/:meeting_id", controllers.UpdateParticipantByUserIdAndMeetingId)

	return e

}
