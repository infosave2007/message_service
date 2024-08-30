package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
)

// AddSupportRatingRequest структура для запроса на добавление оценки техподдержке
type AddSupportRatingRequest struct {
    UserID int `json:"user_id" binding:"required"`
    Rating int `json:"rating" binding:"required"`
    ChatID int `json:"chat_id" binding:"required"`
}

// AverageRatingResponse структура для ответа со средней оценкой
type AverageRatingResponse struct {
    AverageRating float64 `json:"average_rating"`
}

// AddSupportRating добавляет оценку техподдержке
// @Summary Добавление оценки техподдержке
// @Description Этот эндпоинт добавляет новую оценку техподдержке от пользователя.
// @Tags Оценки
// @Accept json
// @Produce json
// @Param rating body AddSupportRatingRequest true "Данные для добавления оценки"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /support/rating [post]
func AddSupportRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var request AddSupportRatingRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    var user models.User
    if err := db.First(&user, request.UserID).Error; err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Пользователь не найден"})
        return
    }

    if request.Rating < 1 || request.Rating > 5 {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Оценка должна быть между 1 и 5"})
        return
    }

    rating := models.SupportRating{
        UserID: request.UserID,
        Rating: request.Rating,
        ChatID: request.ChatID,
    }

    if err := services.SaveSupportRating(&rating, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка сохранения оценки"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Оценка успешно добавлена"})
}

// GetAverageSupportRating возвращает среднюю оценку техподдержки по чату
// @Summary Получение средней оценки техподдержки
// @Description Этот эндпоинт возвращает среднюю оценку техподдержки для указанного чата.
// @Tags Оценки
// @Produce json
// @Param chat_id path int true "ID чата"
// @Success 200 {object} AverageRatingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /support/rating/{chat_id} [get]
func GetAverageSupportRating(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    chatID, err := strconv.Atoi(c.Param("chat_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID чата"})
        return
    }

    avgRating, err := services.GetAverageRatingByChatID(chatID, db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка получения средней оценки"})
        return
    }

    c.JSON(http.StatusOK, AverageRatingResponse{AverageRating: avgRating})
}
