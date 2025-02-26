package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// StartServer запускает сервер с graceful shutdown
func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/about", AboutHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: logMiddleware(mux), // Добавляем логирование
	}

	// Канал для обработки сигналов ОС (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("🚀 Сервер запущен на http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	<-stop
	log.Println("📴 Завершаем работу сервера...")

	// Контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы: %v", err)
	}

	log.Println("✅ Сервер завершил работу.")
}

// logMiddleware логирует запросы
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
