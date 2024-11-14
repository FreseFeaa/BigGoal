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
	Host:         "10.11.21.143",
	Port:         "5672",
	QueueName:    "3",
	ServiceName:  "Fyodr",
	ExchangeName: "it-hub-1",
}
