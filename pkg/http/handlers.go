package http

import (
	"fmt"
	"net/http"
)

func indexHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Todo Server Index\n")
	}
	return http.HandlerFunc(fn)
}

// logRequest logs HTTP requests.
func (s *Server) logRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		s.logger.Infof("HTTP %s %v", req.Method, req.URL)
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}
