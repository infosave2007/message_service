package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "net/http"
)

// RegisterResponse структура для ответа при регистрации пользователя
type RegisterResponse struct {
    Message string `json:"message"`
    UserID  uint   `json:"user_id"`
    Role    uint   `json:"role"`
}

// Register регистрирует нового пользователя
// @Summary Регистрация пользователя
// @Description Этот эндпоинт регистрирует нового пользователя с уникальными именем пользователя и email.
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param user body models.User true "Информация о пользователе"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)  // Получаем объект базы данных из контекста

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    var existingUser models.User
    if err := db.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, ErrorResponse{Error: "Пользователь с таким именем или email уже существует"})
        return
    } else if err != nil && err != gorm.ErrRecordNotFound {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка проверки существования пользователя"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка хеширования пароля"})
        return
    }

    input.Password = string(hashedPassword)

    if err := services.CreateUser(&input, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка создания пользователя"})
        return
    }

    c.JSON(http.StatusOK, RegisterResponse{
        Message: "Пользователь успешно зарегистрирован",
        UserID:  uint(input.ID),    // Преобразование int в uint
        Role:    uint(input.RoleID), // Преобразование int в uint
    })
}

// LoginResponse структура для ответа при аутентификации пользователя
type LoginResponse struct {
    Token string `json:"token"`
    User  struct {
        ID   uint `json:"id"`
        Role uint `json:"role"`
    } `json:"user"`
}

// Login аутентифицирует пользователя
// @Summary Аутентификация пользователя
// @Description Этот эндпоинт выполняет аутентификацию пользователя и возвращает JWT токен.
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param user body models.User true "Информация для аутентификации"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный формат данных"})
        return
    }

    user, err := services.AuthenticateUser(input.Username, input.Password, db)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Неверное имя пользователя или пароль"})
        } else {
            c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка при аутентификации"})
        }
        return
    }

    token, err := services.GenerateToken(uint(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка генерации токена"})
        return
    }

    var response LoginResponse
    response.Token = token
    response.User.ID = uint(user.ID)     // Преобразование int в uint
    response.User.Role = uint(user.RoleID) // Преобразование int в uint

    c.JSON(http.StatusOK, response)
}

// LogoutResponse структура для ответа при выходе пользователя
type LogoutResponse struct {
    Message string `json:"message"`
}

// Logout завершает сессию пользователя
// @Summary Выход пользователя
// @Description Этот эндпоинт завершает текущую сессию пользователя.
// @Tags Пользователи
// @Accept json
// @Produce json
// @Success 200 {object} LogoutResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /logout [post]
func Logout(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    sessionID, exists := c.Get("session_id")
    if !exists {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Сессия не найдена"})
        return
    }

    if err := services.EndSession(sessionID.(uint), db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка завершения сессии"})
        return
    }

    c.JSON(http.StatusOK, LogoutResponse{Message: "Сессия успешно завершена"})
}
