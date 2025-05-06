package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type APIServer struct {
	svc Service
}

func NewAPIServer(svc *Service) *APIServer {
	return &APIServer{
		svc: *svc,
	}
}

func (s *APIServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)

	return http.ListenAndServe(listenAddr, nil)
}

func (s *APIServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetFactFact(context.Background())

	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
