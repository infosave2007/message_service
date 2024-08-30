package middleware

import (
    "message_service/services"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "gorm.io/gorm"
    "net/http"
    "strings"
)

// AuthMiddleware проверяет JWT токен для защищённых маршрутов

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
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Получаем токен из заголовка Authorization
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Необходим токен доступа"})
            c.Abort()
            return
        }

        // Убираем префикс "Bearer " из токена
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Проверка токена
        token, err := services.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный или истекший токен"})
            c.Abort()
            return
        }

        // Извлечение данных из токена (например, user_id)
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
            c.Abort()
            return
        }

        // Передача данных пользователя в контекст для дальнейшего использования
        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}

// DBMiddleware добавляет объект базы данных в контекст запросов
func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    }
}
