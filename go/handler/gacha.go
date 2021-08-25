package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/dbctl"
	"github.com/SuperTikuwa/mission-techdojo/model"
)

func DrawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("x-token")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := dbctl.SelectUserByToken(token)
	if !dbctl.UserExists(user) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var drawRequest model.GachaDrawRequest
	if err := json.Unmarshal(body, &drawRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !dbctl.GachaExists(drawRequest.GachaID) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"success": false,"message":"gacha not found"}`)
		return
	}

	results, err := dbctl.DrawGacha(user, drawRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJson, err := json.Marshal(results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(responseJson))
}
