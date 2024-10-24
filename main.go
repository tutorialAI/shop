package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Создаем новый маршрутизатор
	r := mux.NewRouter()

	// Определяем маршруты
	r.HandleFunc("/auth/login", LoginHandler).Methods("POST")
	r.HandleFunc("/orders/{id:[0-9]+}", OrdersHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", UsersHandler).Methods("GET")

	// Запускаем сервер
	log.Println("API Gateway is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
