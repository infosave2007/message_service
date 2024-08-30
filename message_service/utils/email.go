package utils

import (
    "fmt"
    "net/smtp"
    "strings"
)

// Структура для отправки email сообщений
type Email struct {
    From    string
    To      []string
    Subject string
    Body    string
}

// Функция для отправки email сообщения через SMTP сервер

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
func SendEmail(emailConfig Email, smtpHost, smtpPort, smtpUser, smtpPassword string) error {
    // Формируем сообщение
    msg := "From: " + emailConfig.From + "\n" +
        "To: " + strings.Join(emailConfig.To, ",") + "\n" +
        "Subject: " + emailConfig.Subject + "\n\n" +
        emailConfig.Body

    // Настраиваем аутентификацию SMTP
    auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)

    // Отправляем сообщение через SMTP сервер
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailConfig.From, emailConfig.To, []byte(msg))
    if err != nil {
        return fmt.Errorf("ошибка при отправке email: %v", err)
    }

    return nil
}

// Функция для генерации текста email для уведомлений
func GenerateNotificationEmail(to, subject, body string) Email {
    return Email{
        From:    "noreply@yourvirtualdomain.com",  // Виртуальный домен для отправки
        To:      []string{to},
        Subject: subject,
        Body:    body,
    }
}

// Функция для отправки сообщения через виртуальный email
func SendVirtualEmail(virtualEmail, realEmail, subject, body string, smtpHost, smtpPort, smtpUser, smtpPassword string) error {
    emailConfig := Email{
        From:    virtualEmail,
        To:      []string{realEmail},
        Subject: subject,
        Body:    body,
    }

    return SendEmail(emailConfig, smtpHost, smtpPort, smtpUser, smtpPassword)
}

// Функция для обработки и перенаправления входящего email в чат
func HandleIncomingEmail(virtualEmail, subject, body string, smtpHost, smtpPort, smtpUser, smtpPassword string, realEmail string) error {
    emailConfig := Email{
        From:    virtualEmail,
        To:      []string{realEmail},
        Subject: subject,
        Body:    body,
    }

    return SendEmail(emailConfig, smtpHost, smtpPort, smtpUser, smtpPassword)
}
