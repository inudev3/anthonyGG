package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)
	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.Start(":9090"))
}

type LoggingService struct {
	next Service
}
type CatFact struct {
	Fact string `json:"fact"`
}
type Service interface {
	GetCatFact(ctx context.Context) (*CatFact, error)
}
type CatFactService struct {
	url string
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next}
}
func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v\n", fact, err, time.Since(start))
	}(time.Now())
	fact, err = s.next.GetCatFact(ctx)
	return fact, err
}
func NewCatFactService(url string) Service {
	return &CatFactService{url}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*CatFact, error) {
	res, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	fact := &CatFact{}
	if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
		return nil, err
	}
	return fact, nil
}
