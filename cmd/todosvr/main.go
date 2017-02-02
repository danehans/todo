// Based on: http://thenewstack.io/make-a-restful-json-api-go/
/* curl commands

$ curl http://localhost:8080/todos
$ curl http://localhost:8080/todos/$TODO_ID
$ curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
$ curl -H "Content-Type: application/json" -X "DELETE" http://localhost:8080/todos/3
*/

package main

import (
	"flag"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"

	web "github.com/danehans/todo/pkg/http"
)

var (
	// Defaults to info logging
	log = logrus.New()
)

func main() {
	flags := struct {
		address  string
		logLevel string
	}{}
	flag.StringVar(&flags.address, "address", "127.0.0.1:8080", "HTTP listen address")

	// Log levels https://github.com/Sirupsen/logrus/blob/master/logrus.go#L36
	flag.StringVar(&flags.logLevel, "log-level", "info", "Set the logging level")

	// parse command-line and environment variable arguments
	flag.Parse()

	// validate arguments
	if url, err := url.Parse(flags.address); err != nil || url.String() == "" {
		log.Fatal("A valid HTTP listen address is required")
	}

	// logging setup
	lvl, err := logrus.ParseLevel(flags.logLevel)
	if err != nil {
		log.Fatalf("invalid log-level: %v", err)
	}
	log.Level = lvl

	// HTTP Server
	config := &web.Config{
		Logger: log,
	}
	httpServer := web.NewServer(config)
	log.Infof("Starting Todo HTTP server on %s", flags.address)
	err = http.ListenAndServe(flags.address, httpServer.HTTPHandler())
	if err != nil {
		log.Fatalf("failed to start listening: %v", err)
	}
}
