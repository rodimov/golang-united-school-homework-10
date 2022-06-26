package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["PARAM"]

	_, err := fmt.Fprintf(w, "Hello, %s!", name)

	if err != nil {
		log.Println(err)
	}
}
