package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestNewRouter(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)

	router, _ := NewApp()
	router.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err.Error())
	}
	defer resp.Body.Close()

	want := `{"status": "ok"}`

	if string(got) != want {
		t.Errorf("Expected response body %s, got %s", want, string(got))
	}
}
