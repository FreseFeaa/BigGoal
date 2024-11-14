package api

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен пинг, нужен понг")
	fmt.Fprintf(w, "Ты сигма! 😲😲😲")
	fmt.Println("Понг отправлен")
}

type PingHandler struct {
	requestCount int
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен пинг, нужен понг")
	if h.requestCount >= 5 {
		fmt.Fprintf(w, "Слишком раз много меня хотят(((")
		fmt.Println("Слишком раз много меня хотят(((, requestCount: ", h.requestCount)
		return
	}
	fmt.Fprintf(w, "PONG")
	h.requestCount++
	fmt.Println("Понг отправлен, сейчас requestCount: ", h.requestCount)
}
