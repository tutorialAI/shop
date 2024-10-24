package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var authService *AuthServiceClient

func init() {
	// Создаем клиента для Auth Service
	authService = NewAuthServiceClient("localhost:4000") // Адрес Auth Service
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Обрабатываем запрос на Auth Service
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем Login через Auth Service
	resp, err := authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	// Отправляем JWT токен клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID заказа из URL
	vars := mux.Vars(r)
	orderID := vars["id"]

	// Извлекаем JWT токен из заголовков
	token := r.Header.Get("Authorization")

	// Переадресуем запрос на Order Service
	url := fmt.Sprintf("http://localhost:8082/orders/%s", orderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Unable to create request", http.StatusInternalServerError)
		return
	}

	// Добавляем токен в заголовок запроса
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Unable to reach Order Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Возвращаем ответ клиенту
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из URL
	vars := mux.Vars(r)
	userID := vars["id"]

	// Извлекаем JWT токен из заголовков
	token := r.Header.Get("Authorization")

	// Переадресуем запрос на User Service
	url := fmt.Sprintf("http://localhost:8083/users/%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Unable to create request", http.StatusInternalServerError)
		return
	}

	// Добавляем токен в заголовок запроса
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Unable to reach User Service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Возвращаем ответ клиенту
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
