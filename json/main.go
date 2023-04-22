package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHTTPHandler(getUserById))
	http.ListenAndServe(":3000", nil)
}

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if e, ok := err.(*apiError); !ok {
				writeJson(w, e.Status, e)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{Err: "internal server", Status: http.StatusInternalServerError})
		}
	}
}

type User struct{ ID int }

func getUserById(rw http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Status: http.StatusMethodNotAllowed, Err: "invalid method"}
	}
	return writeJson(rw, http.StatusOK, User{})
}
func writeJson(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
