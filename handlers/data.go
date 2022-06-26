package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
	}

	_, err = fmt.Fprintf(w, "I got message:\n%s", data)

	if err != nil {
		log.Println(err)
	}
}
