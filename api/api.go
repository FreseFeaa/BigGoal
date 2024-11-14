package api

import (
	"fmt"
	"log"
	"net/http"
)

//r.Method  - чтоб узнать метод

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен пинг, нужен понг")
	fmt.Fprintf(w, "pong")
	fmt.Println("Понг отправлен")
}

func handlerApiReceive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ТУТ БУДЕТ КОЛ-ВО Сообщений с типом hello (Получено)")
	// Увеличиваем счётчик
}

func handlerApiSent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ТУТ БУДЕТ КОЛ-ВО Сообщений с типом hello (Отправленно)")
	// Увеличиваем счётчик
}

func Main() {
	fmt.Println("Сервер запущен по: http://localhost:3000")
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/api/v1/receive/messages/hello", handlerApiReceive)
	http.HandleFunc("/api/v1/sent/messages/hello", handlerApiSent)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
