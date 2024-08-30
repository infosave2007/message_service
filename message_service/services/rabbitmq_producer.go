package services

import (
    "github.com/streadway/amqp"
    "message_service/config"
    "log"
    "fmt"
)

// Сервис для отправки сообщений в очередь RabbitMQ

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
func PublishMessageToQueue(message string, cfg *config.Config) error {
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

    q, err := ch.QueueDeclare(
        cfg.RabbitMQQueue, // имя очереди
        true,              // устойчивая очередь
        false,             // автоудаление
        false,             // эксклюзивная очередь
        false,             // без ожидания
        nil,               // аргументы
    )
    if err != nil {
        return fmt.Errorf("не удалось объявить очередь: %v", err)
    }

    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
    if err != nil {
        return fmt.Errorf("не удалось отправить сообщение в очередь: %v", err)
    }

    log.Printf("Сообщение отправлено в очередь: %s", message)
    return nil
}
