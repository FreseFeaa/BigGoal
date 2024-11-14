package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Сервер запущен по: http://localhost:3000")
	http.HandleFunc("/who", PPingHandler)
	http.Handle("/ping2", &PingHandler{})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
