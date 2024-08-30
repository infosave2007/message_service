package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "time"
    "log"
)

// SendMessageRequest структура для запроса отправки сообщения
type SendMessageRequest struct {
    ChatID      int    `json:"chat_id" binding:"required"`
    Content     string `json:"content" binding:"required"`
    UserID      int    `json:"user_id" binding:"required"`
    AttachmentID *int   `json:"attachment_id"`
}

// SendMessageResponse структура для ответа после отправки сообщения
type SendMessageResponse struct {
    Message      string    `json:"message"`
    MessageID    uint      `json:"message_id"`
    Content      string    `json:"content"`
    UserID       int       `json:"user_id"`
    ChatID       int       `json:"chat_id"`
    CreatedAt    time.Time `json:"created_at"`
    AttachmentID *int      `json:"attachment_id,omitempty"`
}

// SendMessage отправляет новое сообщение в чат
// @Summary Отправка нового сообщения
// @Description Этот эндпоинт отправляет новое сообщение в указанный чат.
// @Tags Сообщения
// @Accept json
// @Produce json
// @Param message body SendMessageRequest true "Данные сообщения"
// @Success 200 {object} SendMessageResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /message [post]
func SendMessage(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var request SendMessageRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    if request.UserID <= 0 {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Некорректный ID пользователя"})
        return
    }

    var user models.User
    if err := db.First(&user, request.UserID).Error; err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Пользователь не найден"})
        return
    }

    message := models.Message{
        ChatID:    request.ChatID,
        Content:   request.Content,
        UserID:    request.UserID,
        CreatedAt: time.Now(),
    }

    if request.AttachmentID != nil && *request.AttachmentID > 0 {
        message.AttachedID = request.AttachmentID
    }

    if err := services.SendMessage(&message, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось отправить сообщение"})
        return
    }

    c.JSON(http.StatusOK, SendMessageResponse{
        Message:      "Сообщение успешно отправлено",
        MessageID:    uint(message.ID),  // Преобразуем message.ID в uint
        Content:      message.Content,
        UserID:       message.UserID,
        ChatID:       message.ChatID,
        CreatedAt:    message.CreatedAt,
        AttachmentID: message.AttachedID,
    })
}

// GetMessages возвращает список сообщений в чате
// @Summary Получение списка сообщений
// @Description Этот эндпоинт возвращает список всех сообщений в указанном чате.
// @Tags Сообщения
// @Produce json
// @Param chat_id query int true "ID чата"
// @Success 200 {object} []models.Message
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /messages [get]
func GetMessages(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    chatID, err := strconv.Atoi(c.Query("chat_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    messages, err := services.GetMessagesByChatID(chatID, db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось получить сообщения"})
        return
    }

    c.JSON(http.StatusOK, messages)
}

// MarkMessageAsRead помечает сообщение как прочитанное
// @Summary Пометка сообщения как прочитанного
// @Description Этот эндпоинт помечает сообщение как прочитанное.
// @Tags Сообщения
// @Produce json
// @Param id path int true "ID сообщения"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /message/{id}/read [put]
func MarkMessageAsRead(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    messageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID сообщения"})
        return
    }

    if err := services.MarkMessageAsRead(messageID, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось пометить сообщение как прочитанное"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Сообщение помечено как прочитанное"})
}

// MarkMessageAsChecked помечает сообщение как проверенное
// @Summary Пометка сообщения как проверенного
// @Description Этот эндпоинт помечает сообщение как проверенное.
// @Tags Сообщения
// @Produce json
// @Param id path int true "ID сообщения"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /message/{id}/checked [put]
func MarkMessageAsChecked(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    messageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID сообщения"})
        return
    }

    if err := services.MarkMessageAsChecked(messageID, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось пометить сообщение как проверенное"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Сообщение помечено как проверенное"})
}

// DeleteMessage удаляет сообщение по ID
// @Summary Удаление сообщения
// @Description Этот эндпоинт удаляет сообщение по его ID.
// @Tags Сообщения
// @Produce json
// @Param id path int true "ID сообщения"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /message/{id} [delete]
func DeleteMessage(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    messageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Println("Ошибка: некорректный ID сообщения:", c.Param("id"))
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Некорректный ID сообщения"})
        return
    }

    var message models.Message
    err = db.First(&message, messageID).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            log.Printf("record not found /message/%d\n", messageID)
            c.JSON(http.StatusNotFound, ErrorResponse{Error: "Сообщение с указанным ID не найдено"})
        } else {
            log.Printf("Ошибка при поиске сообщения, ID: %d: %v\n", messageID, err)
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка при поиске сообщения"})
        }
        return
    }

    if err := db.Delete(&message).Error; err != nil {
        log.Printf("Ошибка при удалении сообщения, ID: %d: %v\n", messageID, err)
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось удалить сообщение"})
        return
    }

    log.Printf("Сообщение успешно удалено, ID: %d\n", messageID)
    c.JSON(http.StatusOK, SuccessResponse{Message: "Сообщение успешно удалено"})
}
