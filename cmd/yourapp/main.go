package main

import (
	"github.com/Yessentemir256/forevermatch/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.ListenAndServe(":8080", nil)
}
