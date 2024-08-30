package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "log"
    "time"
)

// CreateChatRequest структура для создания чата
type CreateChatRequest struct {
    Name       string `json:"name" binding:"required"`
    Users      []int  `json:"users" binding:"required"`  // Массив ID пользователей
    AssignedTo int    `json:"assigned_to"`
}

// CreateChatResponse структура для ответа после создания чата
type CreateChatResponse struct {
    Message    string      `json:"message"`
    ChatID     uint        `json:"chat_id"`
    Name       string      `json:"name"`
    Users      []models.User `json:"users"`
    CreatedAt  time.Time   `json:"created_at"`
}

// CreateChat создает новый чат
// @Summary Создание нового чата
// @Description Этот эндпоинт создает новый чат и добавляет пользователей в него.
// @Tags Чаты
// @Accept json
// @Produce json
// @Param chat body CreateChatRequest true "Данные для создания чата"
// @Success 200 {object} CreateChatResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat [post]
func CreateChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input CreateChatRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("Ошибка биндинга JSON: %v", err)
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    var assignedTo *int
    if input.AssignedTo > 0 {
        var assignedUser models.User
        if err := db.First(&assignedUser, input.AssignedTo).Error; err == nil {
            assignedTo = &input.AssignedTo
        }
    }

    chat := models.Chat{
        Name:       input.Name,
        AssignedTo: assignedTo,
        Status:     "active",
    }

    if err := db.Create(&chat).Error; err != nil {
        log.Printf("Ошибка создания чата: %v", err)
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось создать чат"})
        return
    }

    for _, userID := range input.Users {
        chatUser := models.ChatUser{
            ChatID:   chat.ID,
            UserID:   userID,
            JoinedAt: time.Now(),
        }
        if err := db.Create(&chatUser).Error; err != nil {
            log.Printf("Ошибка добавления пользователя с ID %d в чат: %v", userID, err)
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось добавить пользователя в чат"})
            return
        }
    }

    c.JSON(http.StatusOK, CreateChatResponse{
        Message:    "Чат успешно создан",
        ChatID:     uint(chat.ID),
        Name:       chat.Name,
        Users:      chat.Users,
        CreatedAt:  chat.CreatedAt,
    })
}

// GetChat возвращает информацию о чате по ID
// @Summary Получение информации о чате
// @Description Этот эндпоинт возвращает информацию о чате по его ID.
// @Tags Чаты
// @Produce json
// @Param id path int true "ID чата"
// @Success 200 {object} GetChatResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id} [get]
func GetChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    chat, err := services.GetChatByID(chatID, db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось получить информацию о чате"})
        return
    }

    c.JSON(http.StatusOK, GetChatResponse{Chat: *chat}) // Разыменование указателя
}

// GetChatHistoryResponse структура для ответа при получении истории сообщений
type GetChatHistoryResponse struct {
    Messages []models.Message `json:"messages"`
}

// GetChatHistory возвращает историю сообщений в чате
// @Summary Получение истории сообщений в чате
// @Description Этот эндпоинт возвращает историю сообщений в чате по его ID.
// @Tags Чаты
// @Produce json
// @Param id path int true "ID чата"
// @Success 200 {object} GetChatHistoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id}/history [get]
func GetChatHistory(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    var messages []models.Message
    if err := db.Where("chat_id = ?", chatID).Order("created_at asc").Find(&messages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось получить историю сообщений"})
        return
    }

    c.JSON(http.StatusOK, GetChatHistoryResponse{Messages: messages})
}

// AddUserToChatRequest структура для запроса на добавление пользователя в чат
type AddUserToChatRequest struct {
    UserID int `json:"user_id" binding:"required"`
}

// AddUserToChat добавляет пользователя в чат
// @Summary Добавление пользователя в чат
// @Description Этот эндпоинт добавляет пользователя в чат.
// @Tags Чаты
// @Accept json
// @Produce json
// @Param id path int true "ID чата"
// @Param user body AddUserToChatRequest true "Информация о пользователе"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id}/add_user [post]
func AddUserToChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    chat, err := services.GetChatByID(chatID, db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось получить информацию о чате"})
        return
    }

    if chat.Status != "active" {
        c.JSON(http.StatusForbidden, ErrorResponse{Error: "Чат не активен"})
        return
    }

    var input AddUserToChatRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    var user models.User
    if err := db.First(&user, input.UserID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось найти пользователя"})
        return
    }

    if err := services.AddUserToChat(chatID, &user, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось добавить пользователя в чат"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Пользователь успешно добавлен в чат"})
}

// RemoveUserFromChatRequest структура для запроса на удаление пользователя из чата
type RemoveUserFromChatRequest struct {
    UserID int `json:"user_id" binding:"required"`
}

// RemoveUserFromChat удаляет пользователя из чата
// @Summary Удаление пользователя из чата
// @Description Этот эндпоинт удаляет пользователя из чата.
// @Tags Чаты
// @Accept json
// @Produce json
// @Param id path int true "ID чата"
// @Param user body RemoveUserFromChatRequest true "Информация о пользователе"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id}/remove_user [post]
func RemoveUserFromChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    var input RemoveUserFromChatRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    if err := services.RemoveUserFromChat(chatID, &models.User{ID: input.UserID}, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось удалить пользователя из чата"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Пользователь успешно удален из чата"})
}

// CloseChatResponse структура для ответа при закрытии чата
type CloseChatResponse struct {
    Message string `json:"message"`
}

// CloseChat закрывает чат
// @Summary Закрытие чата
// @Description Этот эндпоинт закрывает чат.
// @Tags Чаты
// @Produce json
// @Param id path int true "ID чата"
// @Success 200 {object} CloseChatResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id}/close [post]
func CloseChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    if err := services.CloseChat(chatID, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, CloseChatResponse{Message: "Чат успешно закрыт"})
}

// TransferChatRequest структура для запроса на передачу чата другому пользователю
type TransferChatRequest struct {
    NewUserID int `json:"new_user_id" binding:"required"`
}

// TransferChat передает чат другому пользователю
// @Summary Передача чата другому пользователю
// @Description Этот эндпоинт передает чат другому пользователю.
// @Tags Чаты
// @Accept json
// @Produce json
// @Param id path int true "ID чата"
// @Param transfer body TransferChatRequest true "Данные для передачи чата"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /chat/{id}/transfer [post]
func TransferChat(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    var input TransferChatRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    newUserIDPtr := &input.NewUserID

    if err := services.TransferChat(chatID, newUserIDPtr, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось передать чат"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Чат успешно передан другому пользователю"})
}
