package router

import (

	"os"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gorilla/sessions"
  "github.com/labstack/echo-contrib/session"

	"github.com/MakotoNakai/lets-schedule/controllers"
	"github.com/MakotoNakai/lets-schedule/handlers"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	f, err := os.Create("logs/debug.log")
  if err != nil {
		panic(err)
	}
  defer f.Close()

	req := fmt.Sprintf("Request Body: %s\n", string(reqBody))
	res := fmt.Sprintf("Response Body: %s\n", string(resBody))
	_, err = f.Write([]byte(req))
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(res))
	if err != nil {
		panic(err)
	}
}

func Initialize() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

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
	r.GET("/meetings/user/:user_id", controllers.GetMeetingsByUserId)
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
	r.PUT("/candidate_times/available-time/:meeting_id", controllers.UpdateAvailableTimeByMeetingId)

	r.POST("/participants/new", controllers.CreateParticipant)
	r.GET("/participants/meeting/:meeting_id", controllers.GetParticipantByMeetingId)
	r.PUT("/participants/meeting/:meeting_id", controllers.UpdateParticipantByMeetingId)

	return e

}
