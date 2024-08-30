package services

import (
    "message_service/models"
    "gorm.io/gorm"
)

// Сохранение комнаты техподдержки

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
func SaveSupportRoom(room *models.SupportRoom, db *gorm.DB) error {
    return db.Create(room).Error
}

// Добавление пользователя в комнату техподдержки
func AddUserToRoom(roomID int, user *models.User, db *gorm.DB) error {
    var room models.SupportRoom
    if err := db.First(&room, roomID).Error; err != nil {
        return err
    }

    return db.Model(&room).Association("Users").Append(user)
}
