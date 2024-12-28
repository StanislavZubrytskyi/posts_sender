package main

import (
	"log"
	"net/http"
	"posts_sender/internal/handlers"
	"posts_sender/internal/service"
)

func main() {
	// Ініціалізація сервісів
	dateService := service.NewDateService()
	dateHandler := handlers.NewDateHandler(dateService)

	// Налаштування маршрутів
	http.HandleFunc("/calculate-time", dateHandler.CalculateTimeRemaining)

	// Запуск сервера
	port := ":8080"
	log.Printf("Server is starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
