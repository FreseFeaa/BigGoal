package consumer

import (
	"fmt"
	"log"
	"mb/redis"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Функция для вывода ошибок
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// Создаем структуру Consumer
type Consumer struct {
	UserName    string
	Password    string
	Host        string
	Port        string
	QueueName   string
	ServiceName string
}

// Делаем метод структуры Consumer
func (c *Consumer) Consume() {

	//Подключаемся к RabbitMQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", c.UserName, c.Password, c.Host, c.Port)) //Эти данные лежат в config, передаем их в момент запуска. Для безопастности крч
	failOnError(err, "Не удалось подключиться к RabbitMQ")                                             // Если ошибка при подключении к RabbitMQ
	defer conn.Close()                                                                                 //Гарант, что файл будет закрыт, когда main завершит выполнение, независимо от того, будет ли это нормальный выход или завершение из-за ошибки

	//Создаем канал канал для сообщений
	ch, err := conn.Channel()
	failOnError(err, "Не удалось открыть канал") //Проверка на ошибку при создании канала
	defer ch.Close()                             //гарант закрытия

	msgs, err := ch.Consume(
		c.QueueName, // имя очереди, из которой мы хотим получать сообщения
		"",          // consumerTag: уникальный идентификатор для данного потребителя (можно оставить пустым)
		true,        // autoAck: если true, сообщения будут автоматически подтверждаться после их получения
		false,       // exclusive: если true, очередь будет доступна только для этого потребителя
		false,       // noLocal: если true, потребитель не будет получать сообщения, отправленные им самим
		false,       // noWait: если true, метод не будет ждать подтверждения
		nil,         // аргументы: дополнительные параметры
	)
	failOnError(err, "Не удалось создать ch.Cosume")

	//Делаем канал
	forever := make(chan bool)
	//Запускаем Гоурутину с выводом всех сообщений из Нашего канала msgs
	go func() {
		for d := range msgs {
			log.Printf("- Полученно сообщение: %s\n", d.Body)
			if msgType, ok := d.Headers["type"].(string); ok && msgType == "hello" {
				fmt.Println("Это сообщение с типом: hello")
				redis.Increment("received_hello")
			}

		}
	}()

	//Пишем в консоль всю информацию
	fmt.Println("Успешное подключение к rabbitMq")
	fmt.Println(" [*] - Ожидание сообщений")

	//Блокируем основную гоурутину
	<-forever

	//

	//Немного другой вариант реализации ожидания сообщений, буду тестить
	// for d := range msgs {
	// 	log.Printf("Received a message: %s", d.Body)
	// }

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
