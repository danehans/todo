package http

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Config configures a Server.
type Config struct {
	Logger *logrus.Logger
}

// Server serves Todo's via HTTP.
type Server struct {
	logger *logrus.Logger
}

// NewServer returns a new Server.
func NewServer(config *Config) *Server {
	return &Server{
		logger: config.Logger,
	}
}

// HTTPHandler returns a HTTP handler for the server.
func (s *Server) HTTPHandler() http.Handler {
	mux := http.NewServeMux()

	chain := func(next ContextHandler) http.Handler {
		return s.logRequest(NewHandler(next))
	}

	// server index
	mux.Handle("/", s.logRequest(s.indexHandler()))
	// todo handler
	mux.Handle("/todos", chain(s.todosHandler()))
	mux.Handle("/todos/{todoId}", chain(s.listTodo()))
	return mux
}
