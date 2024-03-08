package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manu-xo5/go-push/api"
)

type apiHandler func(http.ResponseWriter, *http.Request) error

func makeApiHandler(f apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if error := f(w, r); error != nil {
			// handle error
			w.WriteHeader(400)
			fmt.Fprintln(w, error.Error())
		}
	}
}

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeApiHandler(api.UserHandle.Handler))
	router.HandleFunc("/push", makeApiHandler(api.PushHandle.Handler))

	log.Println("JSON api server running")

	http.ListenAndServe(":5000", router)
}
