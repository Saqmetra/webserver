package server

import (
	"fmt"
	"net/http"
)

// HelloHandler отвечает "Hello, World!"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// AboutHandler отвечает информацией о сервере
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/about.html")
}

// LoginPageHandler загружает страницу авторизации
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

// LoginHandler обрабатывает данные формы
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Получаем данные из формы
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Простейшая проверка (для примера)
	if username == "admin" && password == "password" {
		fmt.Fprintln(w, "Авторизация успешна! ✅")
	} else {
		fmt.Fprintln(w, "Ошибка: неверные учетные данные ❌")
	}
}
