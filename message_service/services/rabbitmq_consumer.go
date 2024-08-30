package services

import (
    "github.com/streadway/amqp"
    "message_service/config"
    "gorm.io/gorm"
    "log"
    "fmt"
)

// Сервис для обработки сообщений из очереди RabbitMQ

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
func ConsumeMessagesFromQueue(cfg *config.Config, db *gorm.DB) error {
    conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
        cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort))
    if err != nil {
        return fmt.Errorf("не удалось подключиться к RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        return fmt.Errorf("не удалось открыть канал: %v", err)
    }
    defer ch.Close()

    msgs, err := ch.Consume(
        cfg.RabbitMQQueue, // очередь
        "",                // consumer
        true,              // авто подтверждение
        false,             // эксклюзив
        false,             // без ожидания
        false,             // без аргументов
        nil,               // аргументы
    )
    if err != nil {
        return fmt.Errorf("не удалось получить сообщения из очереди: %v", err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            log.Printf("Получено сообщение: %s", d.Body)

            // Извлечение realEmail на основе virtualEmail
            realEmail, err := GetRealEmailByVirtualEmail(string(d.Body), db)
            if err != nil {
                log.Printf("Ошибка извлечения realEmail: %v", err)
                continue
            }

            // Используем realEmail в вызове HandleIncomingEmail
            err = HandleIncomingEmail(realEmail, "subject", string(d.Body), db, cfg)
            if err != nil {
                log.Printf("Ошибка при обработке сообщения: %v", err)
            }
        }
    }()

    log.Printf(" [*] Ожидание сообщений. Для выхода нажмите CTRL+C")
    <-forever

    return nil
}
