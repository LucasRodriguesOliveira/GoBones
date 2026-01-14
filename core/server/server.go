package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/LucasRodriguesOliveira/GoBones/core/server/pipeline"
	GoBones "github.com/LucasRodriguesOliveira/GoBones/internal/http"
	"github.com/LucasRodriguesOliveira/GoBones/internal/http/exception"
	"github.com/LucasRodriguesOliveira/GoBones/response"
)

type Server struct {
	Config  Config
	Hooks  pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]
	Startup pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]
}

func (s *Server) Start() {
	//TODO: Check if port is in use

	port := fmt.Sprint(s.Config.Port)

	if !isPortAvailable(port) {
		log.Fatalf("Port [%s] not available\n", port)
	}

  if err := s.Startup.Before(nil, nil); err != nil {
    log.Fatal("Could not startup the server...")
  }

  // Normal execution and routes registering
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    req, err := GoBones.ParseRequest(r)
    res := &GoBones.Response{
      Headers: make(map[string]string),
    }

    if err != nil {
      exception.BadRequestException(w)
      return
    }

    if err = s.Hooks.Before(req, res); err != nil {
      exception.BadRequestException(w)
      log.Printf("Error while running `Before` middleware: %v", err)
      return
    }

    // Do Something
    response.OkResponse(res, "Ok")

    if err = s.Hooks.After(req, res); err != nil {
      exception.BadRequestException(w)
      log.Printf("Error while running `After` middleware: %v", err)
      return
    }

    GoBones.WriteResponse(w, res)
  })

  if err := s.Startup.After(nil, nil); err != nil {
    log.Fatal("Error while running stop hooks...")
  }

  if err := http.ListenAndServe(net.JoinHostPort("", port), nil); err != nil {
    log.Fatalln("Could not startup server...", err)
  }
}
