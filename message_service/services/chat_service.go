package services

import (
    "message_service/models"
    "gorm.io/gorm"
)

// Сервис для создания нового чата

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
func CreateChat(chat *models.Chat, db *gorm.DB) error {
    return db.Create(chat).Error
}

// Сервис для получения чата по ID
func GetChatByID(chatID int, db *gorm.DB) (*models.Chat, error) {
    var chat models.Chat
    err := db.Preload("Messages").First(&chat, chatID).Error
    return &chat, err
}

// Сервис для получения истории сообщений чата
func GetChatHistory(chatID int, db *gorm.DB) ([]models.Message, error) {
    return models.GetMessagesByChatID(db, chatID)
}

// Сервис для добавления пользователя в чат
func AddUserToChat(chatID int, user *models.User, db *gorm.DB) error {
    chat, err := GetChatByID(chatID, db)
    if err != nil {
        return err
    }
    return db.Model(&chat).Association("Users").Append(user)
}

// Сервис для удаления пользователя из чата
func RemoveUserFromChat(chatID int, user *models.User, db *gorm.DB) error {
    chat, err := GetChatByID(chatID, db)
    if err != nil {
        return err
    }
    return db.Model(&chat).Association("Users").Delete(user)
}

// Сервис для закрытия чата
func CloseChat(chatID int, db *gorm.DB) error {
    chat, err := GetChatByID(chatID, db)
    if err != nil {
        return err
    }
    chat.Status = "closed"
    return db.Save(chat).Error
}

// Сервис для передачи чата другому пользователю
func TransferChat(chatID int, newUserID *int, db *gorm.DB) error {
    chat, err := GetChatByID(chatID, db)
    if err != nil {
        return err
    }

    // Обновляем поле AssignedTo новым ID пользователя, если указатель не равен nil
    chat.AssignedTo = newUserID
    return db.Save(chat).Error
}
