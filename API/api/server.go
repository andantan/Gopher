package api

import (
	"encoding/json"
	"net/http"
	"opet/API/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handlGetUserByID)
	http.HandleFunc("/user/id", s.handlDeleteUserByID)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handlGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handlDeleteUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}
