package server

import (
	"fmt"
	"net/http"
)

// HelloHandler отвечает "Hello, World!"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, World!")
}

// AboutHandler отвечает информацией о сервере
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Это улучшенный веб-сервер на Go! 🚀")
}
