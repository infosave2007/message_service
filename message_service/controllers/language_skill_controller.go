package controllers

import (
    "message_service/models"
    "message_service/services/language"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

// AddLanguageSkillRequest структура для запроса на добавление языкового навыка
type AddLanguageSkillRequest struct {
    UserID   int    `json:"user_id" binding:"required"`
    Language string `json:"language" binding:"required"`
    Level    string `json:"level" binding:"required"`
}

// AddLanguageSkill добавляет языковой навык сотруднику
// @Summary Добавление языкового навыка
// @Description Этот эндпоинт добавляет новый языковой навык сотруднику.
// @Tags Языковые навыки
// @Accept json
// @Produce json
// @Param skill body AddLanguageSkillRequest true "Данные для добавления языкового навыка"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /language_skill [post]
func AddLanguageSkill(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var skill AddLanguageSkillRequest

    if err := c.ShouldBindJSON(&skill); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    // Проверка на существование языкового навыка для пользователя
    var existingSkill models.LanguageSkill
    if err := db.Where("UserID = ? AND Language = ?", skill.UserID, skill.Language).First(&existingSkill).Error; err == nil {
        c.JSON(http.StatusConflict, ErrorResponse{Error: "Этот языковой навык уже добавлен пользователю"})
        return
    }

    newSkill := models.LanguageSkill{
        UserID:   skill.UserID,
        Language: skill.Language,
        Level:    skill.Level,
    }

    if err := language.SaveLanguageSkill(&newSkill, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка добавления языкового навыка"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Языковой навык успешно добавлен"})
}
