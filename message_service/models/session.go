package models

import (
    "time"
    "gorm.io/gorm"
)

// Модель для таблицы Sessions
type Session struct {
    ID         uint           `gorm:"primaryKey"`
    UserID     uint           `gorm:"column:UserID;not null"`       
    IPAddress  string         `gorm:"column:IPAddress;size:45;not null"` 
    UserAgent  string         `gorm:"column:UserAgent;size:255"`    
    LastActive time.Time      `gorm:"column:LastActive;not null"`   
    CreatedAt  time.Time      `gorm:"column:CreatedAt;not null"`    
    UpdatedAt  time.Time      `gorm:"column:UpdatedAt;not null"`    
    DeletedAt  gorm.DeletedAt `gorm:"column:DeletedAt;index"`       
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
func (Session) TableName() string {
    return "Sessions"
}