package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/dbctl"
	"github.com/SuperTikuwa/mission-techdojo/model"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Method not allowed"}`, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.GenerateToken()

	if err := dbctl.InsertNewUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"token":"`+user.Token+`"}`)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Method not allowed"}`, http.StatusBadRequest)
		return
	}

	token := r.Header[http.CanonicalHeaderKey("x-token")][0]
	user := dbctl.GetUserByToken(token)
	if user.Name == "" {
		http.Error(w, `{"success": false,"message": "User not found"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"name":"`+user.Name+`"}`)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Method not allowed"}`, http.StatusBadRequest)
		return
	}

	token := r.Header[http.CanonicalHeaderKey("x-token")][0]

	if token == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Invalid token"}`, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Token = token

	if dbctl.UserExists(user) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "User not found"}`, http.StatusBadRequest)
		return
	}

	if err := dbctl.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
