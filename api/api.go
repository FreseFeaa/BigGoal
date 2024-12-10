package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Создание пока чт таких переменных
var (
	helloReceivedCount int
	helloSentCount     int
	mu                 sync.Mutex
)

//r.Method  - чтоб узнать метод

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен пинг, нужен понг")
	fmt.Fprintf(w, "pong")
	fmt.Println("Понг отправлен")
}

func handlerApiReceive(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	response := map[string]int{"count": helloReceivedCount}
	fmt.Fprintf(w, "Вот столько сообщений полученно:")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handlerApiSent(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	response := map[string]int{"count": helloSentCount}
	fmt.Fprintf(w, "Вот столько сообщений отправленно:")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func receiveHelloMessage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "Количество полученных сообщений++")
	helloReceivedCount++
}

func Main() {
	fmt.Println("Сервер запущен по: http://localhost:3000")
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/test", receiveHelloMessage)
	http.HandleFunc("/api/v1/receive/messages/hello", handlerApiReceive)
	http.HandleFunc("/api/v1/sent/messages/hello", handlerApiSent)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
