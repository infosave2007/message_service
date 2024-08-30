package models

import (
    "time"
    "gorm.io/gorm"
)

// Модель данных для сообщения
type Message struct {
    ID         int       `json:"id" gorm:"primaryKey"`
    ChatID     int       `json:"chat_id" gorm:"column:ChatID"`
    UserID     int       `json:"user_id" gorm:"column:UserID"`
    Content    string    `json:"content" gorm:"column:Content"`
    CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt"`
    IsRead     bool      `json:"is_read" gorm:"column:IsRead"`
    IsChecked  bool      `json:"is_checked" gorm:"column:IsChecked"`
    AttachedID *int      `json:"attached_id" gorm:"column:AttachedID"`
}


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
func (Message) TableName() string {
    return "Messages"
}

// Создание новой записи сообщения в базе данных
func (m *Message) CreateMessage(db *gorm.DB) error {
    if m.AttachedID == nil || *m.AttachedID == 0 {
        m.AttachedID = nil // Если вложения нет, не связываем с таблицей Attachments
    }
    return db.Create(m).Error
}

// Получение всех сообщений в чате по ID чата
func GetMessagesByChatID(db *gorm.DB, chatID int) ([]Message, error) {
    var messages []Message
    // Получаем все сообщения для указанного чата
    err := db.Where("ChatID = ?", chatID).Order("CreatedAt asc").Find(&messages).Error
    return messages, err
}

// Пометка сообщения как прочитанного
func MarkMessageAsRead(db *gorm.DB, messageID int) error {
    // Обновляем статус сообщения в базе данных
    return db.Model(&Message{}).Where("ID = ?", messageID).Update("IsRead", true).Error
}

// Пометка сообщения как проверенного
func MarkMessageAsChecked(db *gorm.DB, messageID int) error {
    // Обновляем статус сообщения в базе данных
    return db.Model(&Message{}).Where("ID = ?", messageID).Update("IsChecked", true).Error
}
