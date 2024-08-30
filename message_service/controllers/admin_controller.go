package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

// ModerateMessagesResponse структура для ответа при модерации сообщений
type ModerateMessagesResponse struct {
    Message string `json:"message"`
}

// GetRolesResponse структура для ответа с ролями пользователей
type GetRolesResponse struct {
    Roles []string `json:"roles"` // Изменено с []models.Role на []string
}

// UpdateRolesRequest структура для входных данных при обновлении ролей
type UpdateRolesRequest struct {
    UserID  int    `json:"user_id" binding:"required"`
    NewRole string `json:"new_role" binding:"required"`
}

// UpdateRolesResponse структура для ответа при обновлении ролей
type UpdateRolesResponse struct {
    Message string `json:"message"`
}

// GetContainerLogsResponse структура для ответа с логами контейнеров
type GetContainerLogsResponse struct {
    Logs string `json:"logs"` // Изменено с []string на string
}

// CheckContainerStatusResponse структура для ответа со статусом контейнеров
type CheckContainerStatusResponse struct {
    Status string `json:"status"`
}

// UpdateUserResponse структура для ответа при обновлении данных пользователя
type UpdateUserResponse struct {
    Message string `json:"message"`
}

// UpdatePasswordRequest структура для запроса на смену пароля
type UpdatePasswordRequest struct {
    Username    string `json:"username" binding:"required"`
    NewPassword string `json:"new_password" binding:"required"`
}

// UpdatePasswordResponse структура для ответа при смене пароля
type UpdatePasswordResponse struct {
    Message string `json:"message"`
}

// ModerateMessages выполняет модерацию сообщений
// @Summary Модерация сообщений
// @Description Этот эндпоинт выполняет модерацию сообщений.
// @Tags Модерация
// @Accept json
// @Produce json
// @Success 200 {object} ModerateMessagesResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/moderate [get]
func ModerateMessages(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    err := services.ModerateMessages(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка модерации сообщений"})
        return
    }

    c.JSON(http.StatusOK, ModerateMessagesResponse{Message: "Модерация завершена успешно"})
}

// GetRoles возвращает список ролей пользователей
// @Summary Получение ролей пользователей
// @Description Этот эндпоинт возвращает список всех ролей пользователей.
// @Tags Роли
// @Accept json
// @Produce json
// @Success 200 {object} GetRolesResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/roles [get]
func GetRoles(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    roles, err := services.GetRoles(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка получения ролей"})
        return
    }

    c.JSON(http.StatusOK, GetRolesResponse{Roles: roles}) // roles типа []string
}

// UpdateRoles обновляет роли пользователей
// @Summary Обновление ролей пользователей
// @Description Этот эндпоинт обновляет роли пользователей.
// @Tags Роли
// @Accept json
// @Produce json
// @Param roles body UpdateRolesRequest true "Информация о ролях"
// @Success 200 {object} UpdateRolesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/roles [post]
func UpdateRoles(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var rolesInput UpdateRolesRequest
    if err := c.ShouldBindJSON(&rolesInput); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    err := services.UpdateRoles(db, rolesInput.UserID, rolesInput.NewRole)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка обновления ролей"})
        return
    }

    c.JSON(http.StatusOK, UpdateRolesResponse{Message: "Роли обновлены успешно"})
}

// GetContainerLogs возвращает логи контейнеров
// @Summary Получение логов контейнеров
// @Description Этот эндпоинт возвращает логи контейнеров.
// @Tags Контейнеры
// @Accept json
// @Produce json
// @Success 200 {object} GetContainerLogsResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/logs [get]
func GetContainerLogs(c *gin.Context) {
    logs, err := services.GetContainerLogs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка получения логов"})
        return
    }

    c.JSON(http.StatusOK, GetContainerLogsResponse{Logs: logs}) // logs типа string
}

// CheckContainerStatus проверяет статус контейнеров
// @Summary Проверка статуса контейнеров
// @Description Этот эндпоинт проверяет загруженность контейнеров.
// @Tags Контейнеры
// @Accept json
// @Produce json
// @Success 200 {object} CheckContainerStatusResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/status [get]
func CheckContainerStatus(c *gin.Context) {
    status, err := services.CheckContainerStatus()
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка проверки статуса контейнеров"})
        return
    }

    c.JSON(http.StatusOK, CheckContainerStatusResponse{Status: status})
}

// UpdateUser обновляет данные пользователя
// @Summary Обновление данных пользователя
// @Description Этот эндпоинт обновляет информацию о пользователе.
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param user body models.User true "Информация о пользователе"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/update [put]
func UpdateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    if err := services.UpdateUser(&user, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, UpdateUserResponse{Message: "Данные пользователя успешно обновлены"})
}

// UpdatePassword меняет пароль пользователя
// @Summary Смена пароля пользователя
// @Description Этот эндпоинт меняет пароль пользователя.
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param password body UpdatePasswordRequest true "Запрос на смену пароля"
// @Success 200 {object} UpdatePasswordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/update-password [put]
func UpdatePassword(c *gin.Context) {
    var request UpdatePasswordRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    if err := services.UpdateUserPassword(request.Username, request.NewPassword, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, UpdatePasswordResponse{Message: "Пароль успешно обновлен"})
}
