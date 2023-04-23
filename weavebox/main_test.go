package main

import (
	"github.com/twanies/weavebox"
	"net/http"
	"testing"
)

func BenchmarkGetWithValues(b *testing.B) {
	app := weavebox.New()
	app.Get("/hello/:name", func(ctx *weavebox.Context) error { return nil })
	for i := 0; i < b.N; i++ {
		r, err := http.NewRequest(http.MethodGet, "/hello/anthony", nil)
		if err != nil {
			panic(err)
		}
		app.ServeHTTP(nil, r)
	}
}
