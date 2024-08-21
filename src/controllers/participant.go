package controllers

import (
	"log"
	"time"
	"strconv"
	"net/http"
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
    log.Printf("Bind Error: %v", err) // Log the bind error
    return c.JSON(http.StatusBadRequest, err.Error())
  }

  newParticipantList := []models.Participant{}

  for _, ParticipantWithUserName := range newParticipantWithUserNameList {
    Participant := models.Participant{}
    UserName := ParticipantWithUserName.UserName
    Participant.UserId, err = models.GetUserIdFromUserName(db, UserName)
    if err != nil {
      log.Printf("GetUserIdFromUserName Error: %v", err)
      return c.JSON(http.StatusBadRequest, err.Error())
    }

    Participant.MeetingId = ParticipantWithUserName.MeetingId
    Participant.IsHost = ParticipantWithUserName.IsHost
    Participant.HasResponded = ParticipantWithUserName.HasResponded

    Participant.CreatedAt = time.Now()
    Participant.UpdatedAt = time.Now()

    newParticipantList = append(newParticipantList, Participant)
  }

  if models.HostIsInParticipant(newParticipantList) {
    return c.JSON(http.StatusBadRequest, "Host is already a participant")
  }

  result := db.Create(&newParticipantList)
  if result.Error != nil {
    log.Printf("DB Create Error: %v", result.Error)
    return c.JSON(http.StatusBadRequest, result.Error.Error())
  }

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

	pl := models.GetParticipantListByMeetingId(db, mi)
	pwl := models.ConvertToParticipantWithUserNameList(db, pl)

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

	oldParticipantList := models.GetParticipantListByMeetingId(db, mi)
	newParticipantList := []models.Participant{}

	// Start transaction
	tx := db.Begin()

	shorterLength := models.Min(len(oldParticipantList), len(participantWithUserNameList))

	// Process update operations
	for i := 0; i < shorterLength; i++ {
			oldp := oldParticipantList[i]
			pw := participantWithUserNameList[i]
			newp, err := models.ConvertToParticipant(tx, pw) // Use tx in ConvertToParticipant
			if err != nil {
					// If an error occurs, roll back the transaction
					tx.Rollback()
					return c.JSON(http.StatusBadRequest, err.Error())
			}
			// Update participant
			tx.Model(&oldp).Updates(newp)
			newParticipantList = append(newParticipantList, oldp)
	}

	// Delete excess participants if old list is longer
	if len(oldParticipantList) > len(participantWithUserNameList) {
			for i := len(participantWithUserNameList); i < len(oldParticipantList); i++ {
					oldp := oldParticipantList[i]
					if err := tx.Delete(&oldp).Error; err != nil {
							// Rollback and return error
							tx.Rollback()
							return c.JSON(http.StatusBadRequest, err.Error())
					}
			}
	}

	// Create new participants if new list is longer
	if len(oldParticipantList) < len(participantWithUserNameList) {
			for i := len(oldParticipantList); i < len(participantWithUserNameList); i++ {
					pw := participantWithUserNameList[i]
					newp, err := models.ConvertToParticipant(tx, pw) // Use tx in ConvertToParticipant
					if err != nil {
							// Rollback and return error
							tx.Rollback()
							return c.JSON(http.StatusBadRequest, err.Error())
					}
					// Create new participant
					if err := tx.Create(&newp).Error; err != nil {
							// Rollback and return error
							tx.Rollback()
							return c.JSON(http.StatusBadRequest, err.Error())
					}
					newParticipantList = append(newParticipantList, *newp)
			}
	}

	// Commit the transaction if all operations were successful
	if err := tx.Commit().Error; err != nil {
			return c.JSON(http.StatusBadRequest, "Failed to commit transaction")
	}

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
