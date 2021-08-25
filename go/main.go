package main

import (
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/handler"
)

func main() {
	http.HandleFunc("/user/create", handler.CreateHandler)
	http.HandleFunc("/user/get", handler.GetHandler)
	http.HandleFunc("/user/update", handler.UpdateHandler)

	http.HandleFunc("/gacha/draw", handler.DrawHandler)

	http.HandleFunc("/character/list", handler.ListHandler)

	http.HandleFunc("/emission/rate", handler.EmissionRateHandler)

	// http.HandleFunc("/query", handler.Query)

	http.ListenAndServe(":8080", nil)
}
