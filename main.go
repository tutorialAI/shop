package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	name     string
	email    string
	password string
}

var users = []User{
	{name: "Alex", email: "alex@mail.com", password: "123456"},
	{name: "Jon Doe", email: "jon_doe@mail.com", password: "asuehde"},
}

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

func Login(name, email string) (User, bool) {
	for _, user := range users {
		if user.name == name && user.email == email {
			return user, true
		}
	}

	return User{}, false
}
