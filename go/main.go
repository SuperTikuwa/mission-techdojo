package main

import (
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/handler"
)

func main() {
	http.HandleFunc("/user/create", handler.CreateHandler)
	http.HandleFunc("/user/get", handler.GetHandler)
	http.HandleFunc("/user/update", handler.UpdateHandler)
	http.ListenAndServe(":8080", nil)
}
