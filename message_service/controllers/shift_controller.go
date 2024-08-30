package controllers

import (
    "message_service/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "time"
    "strconv"
)

// StartShiftRequest структура для запроса на начало смены
type StartShiftRequest struct {
    UserID int `json:"user_id" binding:"required"`
}

// StartShiftResponse структура для ответа после начала смены
type StartShiftResponse struct {
    Message   string    `json:"message"`
    ShiftID   uint      `json:"shift_id"`
    UserID    int       `json:"user_id"`
    StartTime time.Time `json:"start_time"`
}

// EndShiftResponse структура для ответа после завершения смены
type EndShiftResponse struct {
    Message   string    `json:"message"`
    ShiftID   uint      `json:"shift_id"`
    UserID    int       `json:"user_id"`
    EndTime   time.Time `json:"end_time"`
    Duration  int       `json:"duration"` // Длительность смены в секундах
}

// StartShift начинает смену
// @Summary Начало смены
// @Description Этот эндпоинт начинает новую смену для указанного пользователя.
// @Tags Смены
// @Accept json
// @Produce json
// @Param shift body StartShiftRequest true "Данные для начала смены"
// @Success 200 {object} StartShiftResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shift/start [post]
func StartShift(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var request StartShiftRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    shift := models.Shift{
        UserID:    request.UserID,
        StartTime: time.Now(),
        Active:    true,
    }

    if err := db.Create(&shift).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка начала смены"})
        return
    }

    c.JSON(http.StatusOK, StartShiftResponse{
        Message:   "Смена успешно начата",
        ShiftID:   uint(shift.ID),  // Преобразование shift.ID в uint
        UserID:    shift.UserID,
        StartTime: shift.StartTime,
    })
}

// EndShift завершает смену
// @Summary Завершение смены
// @Description Этот эндпоинт завершает указанную смену.
// @Tags Смены
// @Produce json
// @Param id path int true "ID смены"
// @Success 200 {object} EndShiftResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /shift/end/{id} [post]
func EndShift(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    shiftID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID смены"})
        return
    }

    var shift models.Shift
    if err := db.First(&shift, shiftID).Error; err != nil {
        c.JSON(http.StatusNotFound, ErrorResponse{Error: "Смена не найдена"})
        return
    }

    if shift.StartTime.IsZero() {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Время начала смены не установлено"})
        return
    }

    currentTime := time.Now().Truncate(time.Second)
    shift.EndTime = &currentTime
    shift.Active = false

    if shift.EndTime.Before(shift.StartTime) {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Время окончания смены не может быть раньше времени начала"})
        return
    }

    if err := db.Save(&shift).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка завершения смены"})
        return
    }

    duration := int(shift.EndTime.Sub(shift.StartTime).Seconds())

    c.JSON(http.StatusOK, EndShiftResponse{
        Message:   "Смена успешно завершена",
        ShiftID:   uint(shift.ID),  // Преобразование shift.ID в uint
        UserID:    shift.UserID,
        EndTime:   *shift.EndTime,
        Duration:  duration,
    })
}
