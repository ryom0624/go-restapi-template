package handler

import (
	"log"
	"net/http"
	"webapp/auth"
)

func AuthMiddleware(ac auth.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := ac.FillContext(r)
			if err != nil {
				log.Printf("auth middleware error: %v", err)
				RespondJSON(r.Context(), w, ErrResponse{Error: "Unauthorized"}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func CorsMiddleware(env string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if env == "production" {
				//w.Header().Set("Access-Control-Allow-Origin", "https://exmaple.com")
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*") // TODO: change this to the actual domain
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
