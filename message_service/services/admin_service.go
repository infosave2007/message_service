package services

import (
    "message_service/models"
    "gorm.io/gorm"
    "fmt"
)

// Функция для получения всех ролей пользователей

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
func GetRoles(db *gorm.DB) ([]string, error) {
    var users []models.User
    var roles []string

    // Получаем всех пользователей и загружаем их роли
    if err := db.Preload("Role").Find(&users).Error; err != nil {
        return nil, fmt.Errorf("не удалось получить роли пользователей: %v", err)
    }

    for _, user := range users {
        roles = append(roles, user.Role.Name) // Добавляем имя роли пользователя
    }

    return roles, nil
}

// Функция для обновления роли пользователя
func UpdateRoles(db *gorm.DB, userID int, newRoleName string) error {
    var user models.User
    var role models.Role

    // Находим пользователя по ID
    if err := db.First(&user, userID).Error; err != nil {
        return fmt.Errorf("пользователь не найден: %v", err)
    }

    // Находим роль по имени
    if err := db.Where("name = ?", newRoleName).First(&role).Error; err != nil {
        return fmt.Errorf("роль не найдена: %v", err)
    }

    // Обновляем роль пользователя
    user.RoleID = role.ID
    if err := db.Save(&user).Error; err != nil {
        return fmt.Errorf("не удалось обновить роль пользователя: %v", err)
    }

    return nil
}

// Функция для получения логов контейнеров (заглушка)
func GetContainerLogs() (string, error) {
    // Здесь можно добавить логику для получения реальных логов контейнеров
    return "Логи контейнеров: [здесь будут логи]", nil
}

// Функция для проверки состояния контейнеров (заглушка)
func CheckContainerStatus() (string, error) {
    // Здесь можно добавить логику для проверки реального статуса контейнеров
    return "Статус контейнеров: Все работает корректно", nil
}
