package http

import (
	"context"
	"fmt"
	"net/http"
)

// indexHandler returns a handler that responds with the string "Welcome".
func (s *Server) indexHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome\n")
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

// todoHandler returns a handler that responds with a string "Todo Index".
func (s *Server) todosHandler() ContextHandler {
	fn := func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Todo Index\n")
	}
	return ContextHandlerFunc(fn)
}

// todoHandler returns a handler that responds with a string "Todo Index".
func (s *Server) listTodo() ContextHandler {
	fn := func(ctx context.Context, w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Todo Index\n")
	}
	return ContextHandlerFunc(fn)
}
