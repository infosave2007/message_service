package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
    "fmt"
)

// Структура для хранения JWT ключей
type JWTConfig struct {
    SecretKey string
}

// Структура для JWT токена
type Claims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

// Функция для генерации JWT токена

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
func GenerateJWT(userID int, username string, secretKey string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID:   userID,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", fmt.Errorf("ошибка при генерации JWT токена: %v", err)
    }

    return tokenString, nil
}

// Функция для валидации и разбора JWT токена
func ValidateJWT(tokenString string, secretKey string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return nil, fmt.Errorf("неверная подпись токена")
        }
        return nil, fmt.Errorf("не удалось разобрать токен: %v", err)
    }

    if !token.Valid {
        return nil, fmt.Errorf("недействительный токен")
    }

    return claims, nil
}

// Функция для обновления (refresh) JWT токена
func RefreshJWT(tokenString string, secretKey string) (string, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })

    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            return "", fmt.Errorf("неверная подпись токена")
        }
        return "", fmt.Errorf("не удалось разобрать токен: %v", err)
    }

    if !token.Valid {
        return "", fmt.Errorf("недействительный токен")
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims.ExpiresAt = expirationTime.Unix()

    refreshedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err = refreshedToken.SignedString([]byte(secretKey))
    if err != nil {
        return "", fmt.Errorf("ошибка при обновлении JWT токена: %v", err)
    }

    return tokenString, nil
}
