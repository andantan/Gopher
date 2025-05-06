package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetFactFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetFactFact(ctx context.Context) (*CatFact, error) {
	r, err := http.Get(s.url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	fact := &CatFact{}

	if err := json.NewDecoder(r.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
