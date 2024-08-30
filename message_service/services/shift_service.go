package services

import (
    "message_service/models"
    "gorm.io/gorm"
    "time"
)

// Сохранение смены

// @Summary Краткое описание конечной точки
// @Description Полное описание конечной точки
// @Tags Название тега
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID параметр"
// @Success 200 {object} МодельУспеха
// @Failure 400 {object} МодельОшибки
// @Router /some_endpoint [get]

// Здесь начинается функция

// @Summary Краткое описание конечной точки
// @Description Полное описание конечной точки
// @Tags Название тега
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID параметр"
// @Success 200 {object} МодельУспеха
// @Failure 400 {object} МодельОшибки
// @Router /some_endpoint [get]

// Здесь начинается функция
func SaveShift(shift *models.Shift, db *gorm.DB) error {
    return db.Create(shift).Error
}

// Завершение смены
func EndShift(shiftID int, db *gorm.DB) error {
    var shift models.Shift
    if err := db.First(&shift, shiftID).Error; err != nil {
        return err
    }

    currentTime := time.Now()
    shift.EndTime = &currentTime  // Присваиваем указатель на текущее время
    shift.Active = false

    return db.Save(&shift).Error
}
