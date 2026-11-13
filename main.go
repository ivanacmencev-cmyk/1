package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрируем обработчик для корневого маршрута "/"
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/auth", authHandler) // ← +1 строка

	fmt.Println("Сервер запущен: http://localhost:8080")

	// Запускаем сервер на порту 8080
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Обработчик запросов: принимает writer для ответа и запрос
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет! Это мой первый веб-сервер на Go.")
}

// Новый обработчик: возвращает значение заголовка Authorization
func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header.Get("Authorization"))
}
