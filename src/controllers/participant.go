package controllers

import (

	"time"
	"strconv"
	"net/http"
	"gorm.io/gorm/clause"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

func CreateParticipant(c echo.Context) error {
	
	newParticipantWithUserNameList := []models.ParticipantWithUserName{}
	err := c.Bind(&newParticipantWithUserNameList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newParticipantList := []models.Participant{}

	for _, ParticipantWithUserName := range newParticipantWithUserNameList {

		Participant := models.Participant{}
		UserName := ParticipantWithUserName.UserName
		Participant.UserId = models.GetUserIdFromUserName(UserName)

		Participant.MeetingId = ParticipantWithUserName.MeetingId
		Participant.IsHost = ParticipantWithUserName.IsHost
		Participant.HasResponded = ParticipantWithUserName.HasResponded

		Participant.CreatedAt = time.Now()
		Participant.UpdatedAt = time.Now()

		newParticipantList = append(newParticipantList, Participant)

	}

	db.Create(&newParticipantList)
	return c.JSON(http.StatusCreated, newParticipantList)

}

func GetAllParticipant(c echo.Context) error {

	participantList := []models.Participant{}

	db.Find(&participantList)
	return c.JSON(http.StatusOK, participantList)

}

func GetParticipantByMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pl := models.GetParticipantListByMeetingId(mi)
	pwl := models.ConvertToParticipantWithUserNameList(pl)

	return c.JSON(http.StatusOK, pwl)

}

func UpdateParticipantByUserIdAndMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	participantWithUserNameList := []models.ParticipantWithUserName{}
	err = c.Bind(&participantWithUserNameList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	oldParticipantList := models.GetParticipantListByMeetingId(mi)
	newParticipantList := []models.Participant{}

	db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&models.Participant{})

	tx := db.Begin()

	shorterLength := models.Min(len(oldParticipantList), len(participantWithUserNameList))

	// レコードをUpdateする
	for i := 0; i < shorterLength; i++ {
		oldp := oldParticipantList[i]
		pw := participantWithUserNameList[i]
		newp := models.ConvertToParticipant(pw)
		tx.Model(&oldp).Updates(newp)
		newParticipantList = append(newParticipantList, oldp)
	}

	// 編集前のレコード数が多い場合は、余分なレコードを削除する
	if len(oldParticipantList) > len(participantWithUserNameList) {
		for i := len(participantWithUserNameList); i < len(oldParticipantList); i++ {
			oldp := oldParticipantList[i]
			tx.Delete(&oldp)
		}
	}
	
	// 編集後のレコード数が多い場合は、余ったレコードを新規作成する
	if len(oldParticipantList) < len(participantWithUserNameList) {
		for i := len(oldParticipantList); i < len(participantWithUserNameList); i++ {
			pw := participantWithUserNameList[i]
			newp := models.ConvertToParticipant(pw)
			tx.Create(&newp)
			newParticipantList = append(newParticipantList, newp)
		}
	}
	
	tx.Commit()

	return c.JSON(http.StatusOK, newParticipantList)
}

func DeleteParticipant(c echo.Context) error {

	participant := models.Participant{}
	
	err := c.Bind(&participant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db.Delete(&participant)
	return c.JSON(http.StatusNoContent, participant)

}
