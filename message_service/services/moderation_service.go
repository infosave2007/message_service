package services

import (
    "message_service/models"
    "gorm.io/gorm"
)

// Сервис для модерации сообщений

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
func ModerateMessages(db *gorm.DB) error {
    // Пример логики модерации
    var messages []models.Message

    // Получаем все сообщения, которые еще не были проверены
    if err := db.Where("is_checked = ?", false).Find(&messages).Error; err != nil {
        return err
    }

    for _, message := range messages {
        // Применяем правила модерации
        if containsProhibitedContent(message.Content) {
            // Если сообщение содержит запрещенный контент, удаляем его
            if err := db.Delete(&message).Error; err != nil {
                return err
            }
        } else {
            // Помечаем сообщение как проверенное
            message.IsChecked = true
            if err := db.Save(&message).Error; err != nil {
                return err
            }
        }
    }

    return nil
}

// Пример функции для проверки содержания сообщения
func containsProhibitedContent(content string) bool {
    // Логика проверки содержания сообщения
    // Например, использование ключевых слов или вызов внешнего сервиса для анализа текста
    prohibitedWords := []string{"запрещенное слово 1", "запрещенное слово 2"}

    for _, word := range prohibitedWords {
        if containsWord(content, word) {
            return true
        }
    }
    return false
}

// Функция для проверки наличия слова в тексте
func containsWord(text, word string) bool {
    // Логика поиска слова в тексте
    return len(text) > 0 && len(word) > 0 && (text == word || len(text) > len(word) && (text[:len(word)] == word || text[len(text)-len(word):] == word))
}
