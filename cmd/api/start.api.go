package main

import (
	"mb/api"
	"mb/redis"
)

func main() {
	redis.Main()
	api.Main()
}
