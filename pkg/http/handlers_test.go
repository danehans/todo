package http

import (
	"github.com/Sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

var (
	// Defaults to info logging
	log = logrus.New()
)

func TestHandlers(t *testing.T) {

	tests := []struct {
		name       string
		url        string
		statusCode int
		body       string
	}{
		{"Index", "http://localhost:8080/", http.StatusOK, "Welcome\n"},
		{"Todos", "http://localhost:8080/todos", http.StatusOK, "Todo Index\n"},
	}

	t.Logf("Given the need to test Todo Handlers")

	for _, tt := range tests {

		r, _ := http.NewRequest("GET", tt.url, nil)
		w := httptest.NewRecorder()
		// HTTP Server
		config := &Config{
			Logger: log,
		}
		httpServer := NewServer(config)
		router := httpServer.HTTPHandler()
		t.Logf("\tVerify the %s Get call", tt.name)
		router.ServeHTTP(w, r)

		if w.Code == tt.statusCode {
			t.Logf("\t%s Verify the return code", succeed)
		} else {
			t.Errorf("\t%s Verify the return code", failed)
		}
		if w.Body.String() == tt.body {
			t.Logf("\t%s Verify the return message", succeed)
		} else {
			t.Errorf("\t%s Verify the return message", failed)
		}
	}
}
