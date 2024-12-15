package config

type config struct {
	UserName     string
	Password     string
	Host         string
	Port         string
	QueueName    string
	ServiceName  string
	ExchangeName string
}

var Config config = config{
	UserName:     "guest",
	Password:     "guest",
	Host:         "localhost",
	Port:         "5672",
	QueueName:    "TestQueue",
	ServiceName:  "Fyodr",
	ExchangeName: "it-hub-1",
}
