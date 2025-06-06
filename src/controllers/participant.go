package controllers

import (
	"log"
	"time"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

//----------
// Handlers
//----------

// CreateParticipant creates new participants
// @Summary Create new participants
// @Description Create participants and other details
// @Tags participants
// @Accept json
// @Produce json
// @Param participant body models.ParticipantWithUserName true "List of ParticipantWithUserName"
// @Success 201 {object} models.ParticipantWithUserName
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/participants/new [post]
func CreateParticipant(c echo.Context) error {
  newParticipantWithUserNameList := []models.ParticipantWithUserName{}
  err := c.Bind(&newParticipantWithUserNameList)
  if err != nil {
    log.Printf("Bind Error: %v", err) // Log the bind error
    return c.JSON(http.StatusBadRequest, config.ErrFailedToBindParticipantWithUserNameList)
  }

  newParticipantList := []models.Participant{}

  for _, ParticipantWithUserName := range newParticipantWithUserNameList {
    Participant := models.Participant{}
    UserName := ParticipantWithUserName.UserName
    Participant.UserId, err = models.GetUserIdFromUserName(db, UserName)
    if err != nil {
      return c.JSON(http.StatusBadRequest, config.ErrUserWithUserNameNotFound)
    }

    Participant.MeetingId = ParticipantWithUserName.MeetingId
    Participant.IsHost = ParticipantWithUserName.IsHost
    Participant.HasResponded = ParticipantWithUserName.HasResponded

    Participant.CreatedAt = time.Now()
    Participant.UpdatedAt = time.Now()

    newParticipantList = append(newParticipantList, Participant)
  }

  result := db.Create(&newParticipantList)
  if result.Error != nil {
    log.Printf("DB Create Error: %v", result.Error)
    return c.JSON(http.StatusBadRequest, config.ErrParticipantWithUserNameListFailedToRegister)
  }

  return c.JSON(http.StatusCreated, newParticipantList)
}

// GetParticipantByMeetingId gets participant by meeting id
// @Summary Get participant by meeting id
// @Description Get participant by meeting id
// @Tags participants
// @Accept json
// @Produce json
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.ParticipantWithUserName
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/participants/meeting/{meeting_id} [get]
func GetParticipantByMeetingId(c echo.Context) error {

	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	pl, err := models.GetParticipantListByMeetingId(db, mi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}
	pwl, err := models.ConvertToParticipantWithUserNameList(db, pl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, pwl)

}

// UpdateParticipantByMeetingId updates participant by meeting id
// @Summary Update participant by meeting id
// @Description Update participant by meeting id
// @Tags participants
// @Accept json
// @Produce json
// @Param meeting_id path int true "Meeting ID"
// @Success 200 {object} models.ParticipantWithUserName
// @Failure 400 {object} string "Error message"
// @Router /api/restricted/participants/meeting/{meeting_id} [put]
func UpdateParticipantByMeetingId(c echo.Context) error {
	pmi := c.Param("meeting_id")
	mi, err := strconv.Atoi(pmi)
	if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrIdConversionFailed)
	}

	participantWithUserNameList := []models.ParticipantWithUserName{}
	err = c.Bind(&participantWithUserNameList)
	if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrParticipantWithUserNameListFailedToRegister)
	}

	oldParticipantList, err := models.GetParticipantListByMeetingId(db, mi)
	if err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrMeetingNotFound)
	}

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
					return c.JSON(http.StatusBadRequest, err)
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
							return c.JSON(http.StatusBadRequest, config.ErrFailedToDeleteParticipant)
					}
			}
	}

	// Create new participants if new list is longer
	if len(oldParticipantList) < len(participantWithUserNameList) {
			for i := len(oldParticipantList); i < len(participantWithUserNameList); i++ {
					pw := participantWithUserNameList[i]
					p, err := models.ConvertToParticipant(tx, pw) // Use tx in ConvertToParticipant
					if err != nil {
							// Rollback and return error
							tx.Rollback()
							return c.JSON(http.StatusBadRequest, err)
					}
					// Create new participant
					if err := tx.Create(p).Error; err != nil {
							// Rollback and return error
							tx.Rollback()
							return c.JSON(http.StatusBadRequest, config.ErrFailedToCreateParticipant)
					}
					newParticipantList = append(newParticipantList, p)
			}
	}

	// Commit the transaction if all operations were successful
	if err := tx.Commit().Error; err != nil {
			return c.JSON(http.StatusBadRequest, config.ErrFailedToExecuteTransaction)
	}

	return c.JSON(http.StatusOK, newParticipantList)
}


func DeleteParticipant(c echo.Context) error {

	participant := models.Participant{}
	
	err := c.Bind(&participant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, config.ErrFailedToBindParticipant)
	}

	db.Delete(&participant)
	return c.JSON(http.StatusNoContent, participant)

}
