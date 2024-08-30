package services

import (
    "message_service/models"
    "message_service/config"
    "gorm.io/gorm"
    "fmt"
)

// Сервис для обработки входящих писем и их перенаправления в чат

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
func HandleIncomingEmail(virtualEmail string, subject string, body string, db *gorm.DB, cfg *config.Config) error {
    // Получаем реальный email и идентификатор чата по виртуальному email
    realEmail, chatID, err := GetRealEmailAndChatByVirtual(virtualEmail, db)
    if err != nil {
        return err
    }

    // Создаем новое сообщение в чате
    message := models.Message{
        ChatID:  chatID,
        UserID:  0, // ID системного пользователя
        Content: fmt.Sprintf("Email от %s\nТема: %s\n\n%s", realEmail, subject, body),
    }
    if err := SendMessage(&message, db); err != nil {
        return err
    }

    // Планируем отправку отложенного сообщения, если пользователь не ответит в течение заданного времени
    if err := ScheduleDelayedMessage(chatID, message.Content, cfg.DelayMinutes, db); err != nil {
        return err
    }

    return nil
}

// Сервис для получения реального email и ID чата по виртуальному email
func GetRealEmailAndChatByVirtual(virtualEmail string, db *gorm.DB) (string, int, error) {
    var virtual models.VirtualEmail
    if err := db.Where("virtual_email = ?", virtualEmail).First(&virtual).Error; err != nil {
        return "", 0, fmt.Errorf("виртуальный почтовый ящик не найден: %v", err)
    }
    return virtual.RealEmail, virtual.ChatID, nil
}
// GetRealEmailByVirtualEmail извлекает realEmail на основе virtualEmail
func GetRealEmailByVirtualEmail(virtualEmail string, db *gorm.DB) (string, error) {
    var virtual models.VirtualEmail
    if err := db.Where("virtual_email = ?", virtualEmail).First(&virtual).Error; err != nil {
        return "", fmt.Errorf("виртуальный почтовый ящик не найден: %v", err)
    }
    return virtual.RealEmail, nil
}