package models

import "time"

// Модель данных для комнат техподдержки
type SupportRoom struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    Name      string    `json:"name" gorm:"column:Name;not null"`
    CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt"`
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
func (SupportRoom) TableName() string {
    return "SupportRooms"
}