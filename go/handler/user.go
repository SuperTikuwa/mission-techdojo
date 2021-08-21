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
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var createRequest model.UserCreateRequest
	if err := json.Unmarshal(body, &createRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newUser model.User
	newUser.Name = createRequest.Name
	newUser.GenerateToken()

	if err := dbctl.InsertNewUser(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"token":"`+newUser.Token+`"}`)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Method not allowed"}`, http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-token")

	user := dbctl.SelectUserByToken(token)
	if user.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "User not found"}`, http.StatusBadRequest)
		return
	}

	var getResponse model.UserGetResponse
	getResponse.Name = user.Name

	responseJson, err := json.Marshal(getResponse)
	if err != nil {
		http.Error(w, `{"success": false,"message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(responseJson, getResponse, "hoge")
	fmt.Fprintln(w, string(responseJson))
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"success": false,"message": "Method not allowed"}`, http.StatusBadRequest)
		return
	}

	token := r.Header.Get("x-token")
	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updateRequest model.UserUpdateRequest
	if err := json.Unmarshal(body, &updateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := dbctl.SelectUserByToken(token)
	if !dbctl.UserExists(user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Name = updateRequest.Name

	if err := dbctl.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
