package models

import "time"

// Модель данных для языковых компетенций сотрудников
type LanguageSkill struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    UserID    int       `json:"user_id" gorm:"column:UserID;not null"`
    Language  string    `json:"language" gorm:"column:Language;not null"`
    Level     string    `json:"level" gorm:"column:Level;not null"`
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
func (LanguageSkill) TableName() string {
    return "LanguageSkills"
}