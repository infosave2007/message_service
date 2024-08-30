package services

import (
    "database/sql"  // Добавлен импорт для использования sql.NullFloat64
    "message_service/models"
    "gorm.io/gorm"
    "fmt"
)

// Сохранение оценки техподдержки в базе данных

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
func SaveSupportRating(rating *models.SupportRating, db *gorm.DB) error {
    return db.Create(rating).Error
}

// Получение средней оценки техподдержки по ID чата
func GetAverageRatingByChatID(chatID int, db *gorm.DB) (float64, error) {
    var avgRating sql.NullFloat64
    if err := db.Model(&models.SupportRating{}).Where("ChatID = ?", chatID).Select("AVG(rating)").Scan(&avgRating).Error; err != nil {
        return 0, fmt.Errorf("не удалось получить среднюю оценку: %v", err)
    }

    // Проверяем, является ли значение NULL
    if !avgRating.Valid {
        return 0, nil // Возвращаем 0, если оценок нет
    }

    return avgRating.Float64, nil
}
