package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func HandleHeaders(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))

	if err != nil {
		log.Println(err)
	}

	b, err := strconv.Atoi(r.Header.Get("b"))

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("a+b", strconv.Itoa(a+b))
}
