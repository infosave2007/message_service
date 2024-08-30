package main

import (
    "log"
    "message_service/config"
    "message_service/routes"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "message_service/docs" 
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
    // Загрузка конфигурации
    cfg := config.LoadConfig()

    // Устанавливаем режим работы Gin
    gin.SetMode(cfg.GinMode)

    // Настройка строки подключения DSN
    dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=true"
    log.Printf("Подключение к базе данных с DSN: %s", dsn)

    // Подключение к базе данных с пользовательской стратегией именования
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            NoLowerCase: true, // Отключение автоматического преобразования имен в snake_case
        },
    })
    if err != nil {
        log.Fatalf("Ошибка подключения к базе данных: %v", err)
    }

    // Инициализация Gin маршрутизатора
    r := gin.Default()
    routes.SetupRoutes(r, db)

    // Добавление Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Запуск сервера на указанном порту
    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Ошибка запуска сервера: %v", err)
    }
}
