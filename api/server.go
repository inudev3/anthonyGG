package main

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	listenAddr string
	store      Storage
}

func Newserver(listenaddr string, store Storage) *Server {
	return &Server{listenAddr: listenaddr, store: store}
}
func (s *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)
	json.NewEncoder(w).Encode(user)
}
func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserById)
	return http.ListenAndServe(s.listenAddr, nil)
}
