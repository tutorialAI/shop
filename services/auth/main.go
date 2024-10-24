package main

import (
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("very-secret-key")

func init() {
	makeGRPCServerAndRun()
}

func main() {

}

type User struct {
	name     string
	email    string
	password string
}

var users = []User{
	{name: "Alex", email: "alex@mail.com", password: "123456"},
	{name: "Jon Doe", email: "jon_doe@mail.com", password: "asuehde"},
}

type AuthServer struct {
}

type LoginResponse struct {
	Token   string
	Message string
}

func (s AuthServer) Login(email, password string) LoginResponse {
	for _, user := range users {
		if user.email == email && user.password == password {
			return generateJWTToken(user)
		}
	}

	return LoginResponse{
		Token:   "",
		Message: "User not found",
	}
}

func (s AuthServer) Register() string {
	return ""
}

func generateJWTToken(user User) LoginResponse {
	payload := jwt.MapClaims{
		"sub": user.email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		slog.Error("JWT token signing")
		return LoginResponse{Token: "", Message: "JWT token signing"}
	}

	return LoginResponse{Token: t, Message: "OK"}
}
