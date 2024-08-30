package models

import "time"

// Модель для хранения виртуальных почтовых ящиков и связей с реальными почтовыми ящиками
type VirtualEmail struct {
    ID           int       `json:"id" gorm:"primaryKey;column:ID"`
    RealEmail    string    `json:"real_email" gorm:"column:RealEmail;not null;unique"`
    VirtualEmail string    `json:"virtual_email" gorm:"column:VirtualEmail;not null;unique"`
    ChatID       int       `json:"chat_id" gorm:"column:ChatID;not null"`
    CreatedAt    time.Time `json:"created_at" gorm:"column:CreatedAt"`
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
func (VirtualEmail) TableName() string {
    return "VirtualEmails"
}
// Модель для хранения сообщений, которые необходимо отправить по прошествии времени
type DelayedMessage struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    ChatID    int       `json:"chat_id" gorm:"column:ChatID;not null"`
    Message   string    `json:"message" gorm:"column:Message"`
    SendTime  time.Time `json:"send_time" gorm:"column:SendTime"`
    IsSent    bool      `json:"is_sent" gorm:"column:IsSent"`
    CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
}

// Определяем название таблицы в базе данных
func (DelayedMessage) TableName() string {
    return "DelayedMessages"
}
