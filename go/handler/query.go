package handler

import (
	"net/http"

	"github.com/SuperTikuwa/mission-techdojo/dbctl"
)

func Query(w http.ResponseWriter, r *http.Request) {
	dbctl.Query()
}
