package services

import (
    "message_service/models"
    "message_service/utils"
    "message_service/config"
    "gorm.io/gorm"
    "time"
    "fmt"
    "log"
)

// Сервис для добавления отложенного сообщения

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
func ScheduleDelayedMessage(chatID int, message string, delayMinutes int, db *gorm.DB) error {
    sendTime := time.Now().Add(time.Duration(delayMinutes) * time.Minute)
    delayedMessage := models.DelayedMessage{
        ChatID:   chatID,
        Message:  message,
        SendTime: sendTime,
        IsSent:   false,
    }
    return db.Create(&delayedMessage).Error
}

// Сервис для отправки отложенных сообщений
func ProcessDelayedMessages(db *gorm.DB, cfg *config.Config) error {
    var messages []models.DelayedMessage
    now := time.Now()

    // Получаем все сообщения, время отправки которых уже наступило, но они еще не отправлены
    if err := db.Where("SendTime <= ? AND IsSent = ?", now, false).Find(&messages).Error; err != nil {
        return err
    }

    for _, msg := range messages {
        // Получаем чат
        chat, err := GetChatByID(msg.ChatID, db)
        if err != nil {
            log.Printf("Ошибка при получении чата: %v", err)
            continue
        }

        // Получаем всех пользователей, связанных с чатом через таблицу chat_users
        var chatUsers []models.ChatUser
        if err := db.Where("ChatID = ?", chat.ID).Find(&chatUsers).Error; err != nil {
            log.Printf("Ошибка при получении пользователей чата: %v", err)
            continue
        }

        // Проходим по всем пользователям чата и отправляем email
        for _, chatUser := range chatUsers {
            var user models.User
            if err := db.First(&user, chatUser.UserID).Error; err != nil {
                log.Printf("Ошибка при получении данных пользователя: %v", err)
                continue
            }

            // Генерируем и отправляем email с содержимым чата
            emailBody := fmt.Sprintf("Содержимое чата:\n\n%s", msg.Message)
            emailConfig := utils.GenerateNotificationEmail(user.Email, "Чат - отложенное сообщение", emailBody)

            if err := utils.SendEmail(emailConfig, cfg.EmailServer, cfg.EmailPort, cfg.EmailUser, cfg.EmailPassword); err != nil {
                log.Printf("Ошибка при отправке email: %v", err)
                continue
            }
        }

        // Обновляем статус сообщения как отправленного
        msg.IsSent = true
        if err := db.Save(&msg).Error; err != nil {
            log.Printf("Ошибка при обновлении статуса сообщения: %v", err)
            continue
        }
    }

    return nil
}
