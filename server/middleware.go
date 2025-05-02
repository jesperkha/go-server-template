package server

import (
	"log"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// Add middleware to base handler, in order.
func (s *Server) Add(m ...Middleware) {
	for _, m := range m {
		s.h = m(s.h)
	}
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("")

		h.ServeHTTP(w, r)
	})
}
