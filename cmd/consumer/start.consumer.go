package main

import (
	"mb/config"
	"mb/consumer"
)

func main() {
	con := consumer.Consumer{
		UserName:    config.Config.UserName,
		Password:    config.Config.Password,
		Host:        config.Config.Host,
		Port:        config.Config.Port,
		QueueName:   config.Config.QueueName,
		ServiceName: config.Config.ServiceName,
	}

	con.Consume()
}
