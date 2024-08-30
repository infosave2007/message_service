package config

import (
    "github.com/spf13/viper"
    "log"
)

// Структура для хранения конфигурации приложения
type Config struct {
    ServerPort       string
    DBHost           string
    DBPort           string
    DBUser           string
    DBPassword       string
    DBName           string
    JWTSecret        string
    EmailServer      string
    EmailPort        string
    EmailUser        string
    EmailPassword    string
    VirtualDomain    string // Домен для виртуальных почтовых ящиков
    DelayMinutes     int    // Время задержки перед отправкой содержимого чата
    RabbitMQHost     string // Хост RabbitMQ
    RabbitMQPort     string // Порт RabbitMQ
    RabbitMQUser     string // Пользователь RabbitMQ
    RabbitMQPassword string // Пароль RabbitMQ
    RabbitMQQueue    string // Имя очереди RabbitMQ
    MaxImageSize     int    // Максимальный размер изображения в килобайтах
    MaxDocSize       int    // Максимальный размер документа в мегабайтах
    GinMode          string // Режим работы Gin (debug, release)
}

// Функция для загрузки конфигурации из файла .env

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
func LoadConfig() *Config {
    viper.SetConfigFile(".env")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Ошибка чтения конфигурационного файла: %s", err)
    }

    // Инициализация структуры конфигурации
    config := &Config{
        ServerPort:       viper.GetString("SERVER_PORT"),
        DBHost:           viper.GetString("DB_HOST"),
        DBPort:           viper.GetString("DB_PORT"),
        DBUser:           viper.GetString("DB_USER"),
        DBPassword:       viper.GetString("DB_PASSWORD"),
        DBName:           viper.GetString("DB_NAME"),
        JWTSecret:        viper.GetString("JWT_SECRET"),
        EmailServer:      viper.GetString("EMAIL_SERVER"),
        EmailPort:        viper.GetString("EMAIL_PORT"),
        EmailUser:        viper.GetString("EMAIL_USER"),
        EmailPassword:    viper.GetString("EMAIL_PASSWORD"),
        VirtualDomain:    viper.GetString("VIRTUAL_DOMAIN"),
        DelayMinutes:     viper.GetInt("DELAY_MINUTES"),
        RabbitMQHost:     viper.GetString("RABBITMQ_HOST"),
        RabbitMQPort:     viper.GetString("RABBITMQ_PORT"),
        RabbitMQUser:     viper.GetString("RABBITMQ_USER"),
        RabbitMQPassword: viper.GetString("RABBITMQ_PASSWORD"),
        RabbitMQQueue:    viper.GetString("RABBITMQ_QUEUE"),
        MaxImageSize:     viper.GetInt("MAX_IMAGE_SIZE"),
        MaxDocSize:       viper.GetInt("MAX_DOC_SIZE"),
        GinMode:          viper.GetString("GIN_MODE"), 
    }

    return config
}
