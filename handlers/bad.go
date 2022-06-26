package handlers

import (
	"net/http"
)

func HandleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
