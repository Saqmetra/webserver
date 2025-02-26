package main

import (
	"fmt"
	"net/http"

	// Импортируем пакет server (путь может быть другим, см. пункт 2)
	"webserver/server"
)

func main() {
	http.HandleFunc("/", server.HelloHandler)
	http.HandleFunc("/about", server.AboutHandler)

	// Один маршрут для /login, который обрабатывает и GET, и POST
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			server.LoginPageHandler(w, r) // Отображаем страницу авторизации
		case http.MethodPost:
			server.LoginHandler(w, r) // Обрабатываем вход
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
