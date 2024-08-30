package models

import "time"

// Модель данных для оценки техподдержки
type SupportRating struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    ChatID    int       `json:"chat_id" gorm:"column:ChatID;not null"`
    UserID    int       `json:"user_id" gorm:"column:UserID;not null"`
    Rating    int       `json:"rating" gorm:"column:Rating;not null"`
    Comment   string    `json:"comment" gorm:"column:Comment"`
    CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
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
func (SupportRating) TableName() string {
    return "SupportRatings"
}