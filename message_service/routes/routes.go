package routes

import (
    "github.com/gin-gonic/gin"
    "message_service/controllers"
    "message_service/middleware" 
    "gorm.io/gorm"
)

// Функция для установки маршрутов

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
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    // Добавляем middleware для передачи базы данных в контексте
    r.Use(middleware.DBMiddleware(db))
    
    // Маршруты для аутентификации и регистрации
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.PUT("/user/update", controllers.UpdateUser)
    r.PUT("/user/update-password", controllers.UpdatePassword)
    r.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
    
    // Защищённые маршруты
    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware()) // Применение middleware для проверки токенов

    // Маршруты для работы с сообщениями
    protected.GET("/messages", controllers.GetMessages)
    protected.POST("/message", controllers.SendMessage)
    protected.PUT("/message/:id/read", controllers.MarkMessageAsRead)
    protected.DELETE("/message/:id", controllers.DeleteMessage)

    // Маршруты для работы с вложениями
    protected.POST("/attachment", controllers.UploadAttachment)
    protected.GET("/attachment/:id", controllers.DownloadAttachment)

    // Маршруты для работы с чатами
    protected.GET("/chat/:id/history", controllers.GetChatHistory)
    protected.POST("/chat", controllers.CreateChat)
    protected.GET("/chat/:id", controllers.GetChat)
    protected.POST("/chat/:id/add_user", controllers.AddUserToChat)
    protected.POST("/chat/:id/remove_user", controllers.RemoveUserFromChat)
    protected.POST("/chat/:id/transfer", controllers.TransferChat)
    protected.POST("/chat/:id/close", controllers.CloseChat)

    // Маршруты для работы с техподдержкой и оценками
    protected.POST("/support/rating", controllers.AddSupportRating)
    protected.GET("/support/rating/:chat_id", controllers.GetAverageSupportRating)

    // Маршруты для управления сменами
    protected.POST("/shift/start", controllers.StartShift)
    protected.POST("/shift/end/:id", controllers.EndShift)

    // Маршруты для работы с комнатами техподдержки
    protected.POST("/support_room", controllers.CreateSupportRoom)
    protected.POST("/support_room/:id/add_user", controllers.AddUserToSupportRoom)
    // Маршрут для получения статуса комнаты техподдержки
    protected.GET("/support_room/:id/status",controllers.GetSupportRoomStatus)

    // Маршруты для работы с языковыми навыками сотрудников
    protected.POST("/language_skill", controllers.AddLanguageSkill)

    // Административные маршруты (также защищённые)
    protected.GET("/admin/moderate", controllers.ModerateMessages)
    protected.GET("/admin/roles", controllers.GetRoles)
    protected.POST("/admin/roles", controllers.UpdateRoles)
    protected.GET("/admin/logs", controllers.GetContainerLogs)
    protected.GET("/admin/status", controllers.CheckContainerStatus)
}
