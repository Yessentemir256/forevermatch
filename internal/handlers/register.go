package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Прочитать данные из тела POST запроса
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
			return
		}

		// Получить данные пользователя из POST запроса
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Вывести информацию о регистрации пользователя с помощью пакета log
		log.Printf("Пользователь %s с email %s и паролем %s успешно зарегистрирован", username, email, password)

		// Отправить ответ клиенту
		fmt.Fprintf(w, "Пользователь %s успешно зарегистрирован", username)
	} else {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}
