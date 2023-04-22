package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFooRR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handleGetFoo(rr, req)
	res := rr.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", res.StatusCode)
	}
	expected := "foo"
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", res.StatusCode)
	}
	expected := "foo"
	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
