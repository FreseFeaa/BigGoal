package producer

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Функция для вывода ошибок
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Создаем структуру Producer
type Producer struct {
	UserName      string
	Password      string
	Host          string
	Port          string
	QueueNameSent string
	ServiceName   string
	ExchangeName  string
}

// Делаем метод структуры Producer
func (p *Producer) Produce(routingKey, messageType, body string) { //Передаём в него - routingKey, messageType, body
	//Подключаемся к RabbitMQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", p.UserName, p.Password, p.Host, p.Port)) //Эти данные лежат в config, передаем их в момент запуска. Для безопастности крч
	failOnError(err, "Не удалось подключиться к RabbitMQ")                                             // Если ошибка при подключении к RabbitMQ
	defer conn.Close()

	//Создаем канал канал для сообщений
	ch, err := conn.Channel()
	failOnError(err, "Не удалось открыть канал") //Проверка на ошибку при создании канала
	defer ch.Close()                             //гарант закрытия

	//Объявление очереди
	q, err := ch.QueueDeclare(
		p.QueueNameSent,
		true,  // durable: очередь не будет сохраняться при перезапуске сервера
		false, // autoDelete: очередь не будет удалена, когда все подписчики отключатся
		false, // exclusive: очередь не будет эксклюзивной для данного соединения
		false, // noWait: не ждать подтверждения создания очереди
		nil,   // аргументы (можно передать дополнительные параметры)
	)
	failOnError(err, "He удалось создать/подключиться к очереди")
	fmt.Println("Информация об очереди: ", q) //Выводим информацию об очереди
	// контекст с таймаутом, который используется для управления временем выполнения операций
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"Meow",     // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: messageType,
			Headers: amqp.Table{
				"type": "hello",
			},
			Body: []byte(body),
		})
	failOnError(err, "Не удалось отправить сообщение(")

	log.Printf(" - Отправлено: %s\n", body)
}
