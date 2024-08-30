package controllers
import (
    "message_service/models"
)
// ErrorResponse структура для ответа при ошибке
type ErrorResponse struct {
    Error string `json:"error"`
}
// SuccessResponse структура для общего успешного ответа
type SuccessResponse struct {
    Message string `json:"message"`
}
// GetChatResponse структура для ответа при получении информации о чате
type GetChatResponse struct {
    Chat models.Chat `json:"chat"`
}