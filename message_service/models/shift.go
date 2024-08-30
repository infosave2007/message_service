package models

import "time"

// Модель данных для управления сменами
type Shift struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    UserID    int       `json:"user_id" gorm:"column:UserID;not null"`       // ID сотрудника
    StartTime time.Time `json:"start_time" gorm:"column:StartTime;not null"` // Начало смены
    EndTime   *time.Time `json:"end_time" gorm:"column:EndTime"`             // Конец смены, может быть NULL
    Active    bool      `json:"active" gorm:"column:Active;default:true"`    // Активность смены
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
func (Shift) TableName() string {
    return "Shifts"
}
