package services

import (
    "message_service/models"
    "gorm.io/gorm"
)

// Сервис для отправки нового сообщения в чат

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
func SendMessage(message *models.Message, db *gorm.DB) error {
    return message.CreateMessage(db)
}

// Сервис для получения сообщений по ID чата
func GetMessagesByChatID(chatID int, db *gorm.DB) ([]models.Message, error) {
    return models.GetMessagesByChatID(db, chatID)
}

// Сервис для пометки сообщения как прочитанного
func MarkMessageAsRead(messageID int, db *gorm.DB) error {
    return models.MarkMessageAsRead(db, messageID)
}

// Сервис для пометки сообщения как проверенного
func MarkMessageAsChecked(messageID int, db *gorm.DB) error {
    return models.MarkMessageAsChecked(db, messageID)
}
