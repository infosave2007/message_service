package language

import (
    "message_service/models"
    "gorm.io/gorm"
    "fmt"
)

// SaveLanguageSkill сохраняет языковой навык сотрудника

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
func SaveLanguageSkill(skill *models.LanguageSkill, db *gorm.DB) error {
    if err := db.Create(skill).Error; err != nil {
        return fmt.Errorf("не удалось сохранить языковой навык: %v", err)
    }
    return nil
}
