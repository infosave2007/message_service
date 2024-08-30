package controllers

import (
    "message_service/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "time"
)

// CreateSupportRoomRequest структура для запроса на создание комнаты техподдержки
type CreateSupportRoomRequest struct {
    Name  string `json:"name" binding:"required"`
    Users []int  `json:"users" binding:"required"` // Массив ID пользователей
}

// AddUserToSupportRoomRequest структура для запроса на добавление пользователя в комнату
type AddUserToSupportRoomRequest struct {
    UserID int `json:"user_id" binding:"required"`
}

// CreateSupportRoomResponse структура для ответа после создания комнаты техподдержки
type CreateSupportRoomResponse struct {
    Message   string    `json:"message"`
    RoomID    uint      `json:"room_id"`
    Name      string    `json:"name"`
    Users     []int     `json:"users"`
    CreatedAt time.Time `json:"created_at"`
}

// UserDTO структура для информации о пользователе в комнате
type UserDTO struct {
    ID   int    `json:"id"`
    Role string `json:"role"`
}

// GetSupportRoomStatusResponse структура для ответа о статусе комнаты техподдержки
type GetSupportRoomStatusResponse struct {
    Message string    `json:"message"`
    RoomID  uint      `json:"room_id"`
    Name    string    `json:"name"`
    Users   []UserDTO `json:"users"`
    Status  string    `json:"status"`
}

// CreateSupportRoom создает комнату техподдержки
// @Summary Создание комнаты техподдержки
// @Description Этот эндпоинт создает новую комнату техподдержки и добавляет в нее указанных пользователей.
// @Tags Комнаты техподдержки
// @Accept json
// @Produce json
// @Param room body CreateSupportRoomRequest true "Данные для создания комнаты техподдержки"
// @Success 200 {object} CreateSupportRoomResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /support_room [post]
func CreateSupportRoom(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input CreateSupportRoomRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    room := models.Chat{
        Name:      input.Name,
        Status:    "support_room",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := db.Create(&room).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка создания комнаты"})
        return
    }

    for _, userID := range input.Users {
        chatUser := models.ChatUser{
            ChatID:   room.ID,
            UserID:   userID,
            JoinedAt: time.Now(),
        }
        if err := db.Create(&chatUser).Error; err != nil {
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка добавления пользователя в комнату"})
            return
        }
    }

    c.JSON(http.StatusOK, CreateSupportRoomResponse{
        Message:   "Комната успешно создана",
        RoomID:    uint(room.ID),  // Преобразование room.ID в uint
        Name:      room.Name,
        Users:     input.Users,
        CreatedAt: room.CreatedAt,
    })
}

// AddUserToSupportRoom добавляет пользователя в комнату техподдержки
// @Summary Добавление пользователя в комнату техподдержки
// @Description Этот эндпоинт добавляет пользователя в указанную комнату техподдержки.
// @Tags Комнаты техподдержки
// @Accept json
// @Produce json
// @Param id path int true "ID комнаты техподдержки"
// @Param user body AddUserToSupportRoomRequest true "Данные для добавления пользователя"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /support_room/{id}/add_user [post]
func AddUserToSupportRoom(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    roomID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID комнаты"})
        return
    }

    var input AddUserToSupportRoomRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных или некорректный UserID"})
        return
    }

    var existingChatUser models.ChatUser
    if err := db.Where("ChatID = ? AND UserID = ?", roomID, input.UserID).First(&existingChatUser).Error; err == nil {
        c.JSON(http.StatusConflict, ErrorResponse{Error: "Пользователь уже находится в комнате"})
        return
    }

    chatUser := models.ChatUser{
        ChatID:   roomID,
        UserID:   input.UserID,
        JoinedAt: time.Now(),
    }

    if err := db.Create(&chatUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка добавления пользователя в комнату"})
        return
    }

    c.JSON(http.StatusOK, SuccessResponse{Message: "Пользователь успешно добавлен в комнату"})
}

// GetSupportRoomStatus возвращает статус комнаты техподдержки
// @Summary Получение статуса комнаты техподдержки
// @Description Этот эндпоинт возвращает статус указанной комнаты техподдержки, включая список пользователей.
// @Tags Комнаты техподдержки
// @Produce json
// @Param id path int true "ID комнаты техподдержки"
// @Success 200 {object} GetSupportRoomStatusResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /support_room/{id}/status [get]
func GetSupportRoomStatus(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    roomID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID комнаты"})
        return
    }

    var room models.Chat
    if err := db.First(&room, roomID).Error; err != nil {
        c.JSON(http.StatusNotFound, ErrorResponse{Error: "Комната не найдена"})
        return
    }

    if room.Status != "support_room" {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Это не комната техподдержки"})
        return
    }

    var users []models.User
    if err := db.Model(&room).Association("Users").Find(&users); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось получить пользователей комнаты"})
        return
    }

    userDTOs := make([]UserDTO, len(users))
    for i, user := range users {
        var role models.Role
        db.First(&role, user.RoleID)

        userDTOs[i] = UserDTO{
            ID:   user.ID,
            Role: role.Name,
        }
    }

    c.JSON(http.StatusOK, GetSupportRoomStatusResponse{
        Message: "Статус комнаты успешно получен",
        RoomID:  uint(room.ID),  // Преобразование room.ID в uint
        Name:    room.Name,
        Users:   userDTOs,
        Status:  room.Status,
    })
}
