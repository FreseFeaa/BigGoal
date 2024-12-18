package api

import (
	"encoding/json"
	"fmt"
	"log"
	"mb/redis"
	"net/http"
	"sync"
)

var mu sync.Mutex

// Блок для тестов
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен пинг, нужен понг")
	fmt.Fprintf(w, "pong")
	fmt.Println("Понг отправлен")
}

func receiveHelloMessageTest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "Количество полученных сообщений++")
	redis.Increment("received_hello")
}

func SentHelloMessageTest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "Количество полученных сообщений++")
	redis.Increment("sent_hello")
}

// Блок по заданию
func handlerApiReceive(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count, _ := redis.GetReceivedHelloCount()
	response := map[string]int{"count": int(count)}
	fmt.Fprintf(w, "Вот столько сообщений полученно:")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handlerApiSent(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count, _ := redis.GetSentHelloCount()
	response := map[string]int{"count": int(count)}
	fmt.Fprintf(w, "Вот столько сообщений отправленно:")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Main() {

	//Блок для тестов
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/receive", receiveHelloMessageTest)
	http.HandleFunc("/sent", SentHelloMessageTest)

	//Блок по заданию
	http.HandleFunc("/api/v1/receive/messages/hello", handlerApiReceive)
	http.HandleFunc("/api/v1/sent/messages/hello", handlerApiSent)

	//Запуск сервера
	fmt.Println("Сервер запущен по: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
