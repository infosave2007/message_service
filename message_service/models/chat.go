package models

import "time"


// Модель данных для чата
type Chat struct {
    ID         int       `json:"id" gorm:"primaryKey;column:ID"`
    Name       string    `json:"name" gorm:"column:Name"`
    CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt"`
    UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt"`
    UserIDs    []int     `json:"users" gorm:"-"`  // Массив ID пользователей, не сохраняем в базе напрямую
    AssignedTo *int      `json:"assigned_to" gorm:"column:AssignedTo"`
    EntryPoint string    `json:"entry_point" gorm:"column:EntryPoint"`
    Status     string    `json:"status" gorm:"column:Status;default:active"`
    Users      []User    `json:"users" gorm:"many2many:ChatUsers;joinForeignKey:ChatID;joinReferences:UserID"`
    Messages   []Message `json:"messages" gorm:"foreignKey:ChatID"`
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
func (Chat) TableName() string {
    return "Chats"
}
// Структура для связи между чатами и пользователями (многие ко многим)
    type ChatUser struct {
        ChatID   int       `gorm:"primaryKey;column:ChatID"`   // Поле ChatID соответствует столбцу ChatID в базе данных
        UserID   int       `gorm:"primaryKey;column:UserID"`   // Поле UserID соответствует столбцу UserID в базе данных
        JoinedAt time.Time `gorm:"column:JoinedAt"`            // Поле JoinedAt соответствует столбцу JoinedAt в базе данных
    }
// TableName указывает GORM использовать имя таблицы с учетом регистра
func (ChatUser) TableName() string {
    return "ChatUsers"
}