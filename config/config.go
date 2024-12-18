package config

type config struct {
	UserName      string
	Password      string
	Host          string
	Port          string
	QueueName     string
	QueueNameSent string
	ServiceName   string
	ExchangeName  string
}

var Config config = config{
	UserName:      "guest",
	Password:      "guest",
	Host:          "localhost",
	Port:          "5672",
	QueueName:     "TestQueue",
	QueueNameSent: "TestSentQueue",
	ServiceName:   "Fyodr",
	ExchangeName:  "Meow",
}
