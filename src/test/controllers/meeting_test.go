package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MakotoNakai/lets-schedule/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateMeetingPhysicalSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"物理開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingPhysicalButNoTitle(t *testing.T) {

	meetingJSON := `{"title":"","description":"hoge","type":"物理開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingPhysicalButNoPlace(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"物理開催","place":"", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingOnlineSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"オンライン開催","place":"", "url":"https://hoge.com", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingOnlineButNoURL(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"オンライン開催","place":"", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingHybridSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"https://hoge.com", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingHybridButNoPlace(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"", "url":"https://hoge.com", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestCreateMeetingHybridButNoURL(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/meetings/new", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetMeetingByMeetingNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetMeetingById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForHostNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForHostNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForHostNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForHost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetConfirmedMeetingsForGuestNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotConfirmedMeetingsForGuestNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotConfirmedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestSuccess(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestEmptyID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/''", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestHugeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/100", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestZeroID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/0", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestNegativeID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/-1", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}

func TestGetNotRespondedMeetingsForGuestNullID(t *testing.T) {

	meetingJSON := `{"title":"hoge","description":"hoge","type":"ハイブリッド開催","place":"SFC", "url":"", "all_participants_responded":false, "is_confirmed":false, "start_time":"2022-09-19 22:00:00","end_time":"2022-09-26 22:00:00", "hour":1}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/meetings/", strings.NewReader(meetingJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetNotRespondedMeetingsForGuest(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, meetingJSON, rec.Body.String())
	}
}
