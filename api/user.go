package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userHandler struct{}

type User struct {
	Token string `json:"token"`
}

var data = []User{}

// method handler
func (p *userHandler) Handler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return p.get(w, r)
	}

	if r.Method == "POST" {
		return p.post(w, r)
	}

	return fmt.Errorf("method not allowed")
}

// get subscriptions
func (p *userHandler) get(w http.ResponseWriter, r *http.Request) error {
	println(r.RequestURI)

	data, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Write(data)
	return nil
}

// save subscription
func (p *userHandler) post(w http.ResponseWriter, r *http.Request) error {
	newSub := User{}

	if err := json.NewDecoder(r.Body).Decode(&newSub); err != nil {
		return err
	}

	data = append(data, newSub)

	fmt.Fprintln(w, "saved user")

	return nil
}

var UserHandle = userHandler{}
