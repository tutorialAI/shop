package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Обрабатываем запрос на Auth Service
	resp, err := http.Post("http://localhost:8081/auth/login", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Unable to authenticate", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от Auth Service
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read response", http.StatusInternalServerError)
		return
	}

	// Возвращаем клиенту результат
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
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
