package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/dbctl"
	"github.com/SuperTikuwa/mission-techdojo/model"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("x-token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := dbctl.SelectUserByToken(token)
	if user.Token != token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userCharacter, err := dbctl.SelectAllUserCharacter(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	listResponse := model.CharacterListResponse{
		Characters: userCharacter,
	}

	responseJson, err := json.Marshal(listResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(responseJson))
}
