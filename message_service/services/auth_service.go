package services

import (
    "time"
    "errors"
    "message_service/models"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// Секретный ключ для подписи JWT токенов
var jwtSecret = []byte("ваш_секретный_ключ")

// HashPassword хеширует пароль с использованием bcrypt

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
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Сервис для создания нового пользователя
func CreateUser(user *models.User, db *gorm.DB) error {
    return user.CreateUser(db)
}

// Сервис для получения пользователя по имени пользователя (username)
func GetUserByUsername(username string, db *gorm.DB) (*models.User, error) {
    return models.GetUserByUsername(db, username)
}

// Сервис для обновления данных пользователя
func UpdateUser(user *models.User, db *gorm.DB) error {
    return user.UpdateUser(db)
}

// Сервис для смены пароля пользователя
func UpdateUserPassword(username, newPassword string, db *gorm.DB) error {
    user, err := models.GetUserByUsername(db, username)
    if err != nil {
        return err
    }

    // Обновляем пароль
    hashedPassword, err := HashPassword(newPassword)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    return user.UpdateUser(db)
}

// Сервис для удаления пользователя
func DeleteUser(id int, db *gorm.DB) error {
    return models.DeleteUser(db, id)
}

// Сервис для проверки пользователя и пароля (аутентификация)
func AuthenticateUser(username, password string, db *gorm.DB) (*models.User, error) {
    user, err := GetUserByUsername(username, db)
    if err != nil {
        return nil, err
    }

    // Проверяем хеш пароля
    if err := models.ComparePassword(user.Password, password); err != nil {
        return nil, err
    }

    return user, nil
}

// Генерация JWT токена
func GenerateToken(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Токен истекает через 24 часа
    })
    return token.SignedString(jwtSecret)
}

// Проверка JWT токена
func ValidateToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("Неверный или истекший токен")
    }
    return token, nil
}

// Создание сессии для пользователя
func CreateSession(userID uint, ipAddress, userAgent string, db *gorm.DB) error {
    session := &models.Session{
        UserID:     userID,
        IPAddress:  ipAddress,
        UserAgent:  userAgent,
        LastActive: time.Now(),
        CreatedAt:  time.Now(),
        UpdatedAt:  time.Now(),
    }
    return db.Create(session).Error
}

// Обновление активности сессии
func UpdateSessionActivity(sessionID uint, db *gorm.DB) error {
    return db.Model(&models.Session{}).Where("id = ?", sessionID).Update("last_active", time.Now()).Error
}

// Завершение сессии (Logout)
func EndSession(sessionID uint, db *gorm.DB) error {
    return db.Delete(&models.Session{}, sessionID).Error
}

// Сервис для получения роли пользователя по ID
func GetRoleByID(roleID int, db *gorm.DB) (*models.Role, error) {
    var role models.Role
    err := db.First(&role, roleID).Error
    return &role, err
}

// Сервис для получения всех ролей
func GetAllRoles(db *gorm.DB) ([]models.Role, error) {
    var roles []models.Role
    err := db.Find(&roles).Error
    return roles, err
}
