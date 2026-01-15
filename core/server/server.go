package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/LucasRodriguesOliveira/GoBones/core/pipeline"
	"github.com/LucasRodriguesOliveira/GoBones/core/router"
	GoBones "github.com/LucasRodriguesOliveira/GoBones/internal/http"
	"github.com/LucasRodriguesOliveira/GoBones/internal/http/exception"
)

type Server struct {
	Config  Config
	Hooks   pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]
	Startup pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]
	Router  router.Router
}

func New(port int) *Server {
	return &Server{
		Config:  Config{Port: port},
		Startup: pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]{},
		Hooks:   pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]{},
		Router: *router.New(),
	}
}

func (s *Server) Start() {
	//TODO: Check if port is in use

	port := fmt.Sprint(s.Config.Port)

	if isPortInUse(port) {
		log.Fatalf("Port [%s] not available\n", port)
	}

	if err := s.Startup.Before(nil, nil); err != nil {
		log.Fatal("Error while running startup `before` hooks...")
	}

	// Normal execution and routes registering
	for path, handlers := range s.Router.GetRoutes() {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()

			req, err := GoBones.ParseRequest(r)
			res := &GoBones.Response{
				Headers: make(map[string]string),
			}

			if err != nil {
				exception.BadRequestException(w)
				return
			}

			handler, ok := handlers[req.Method]

			if !ok {
				exception.BadRequestException(w)
			}

			if err = s.Hooks.Before(req, res); err != nil {
				exception.BadRequestException(w)
				log.Printf("Error while running `Before` middleware: %v", err)
				return
			}

			// Do Something
			err = handler(req, res)

			if err != nil {
				log.Printf("Error during execution: %v", err)
			}

			if err = s.Hooks.After(req, res); err != nil {
				exception.BadRequestException(w)
				log.Printf("Error while running `After` middleware: %v", err)
				return
			}

			GoBones.WriteResponse(w, res)
		})
	}

	if err := s.Startup.After(nil, nil); err != nil {
		log.Fatal("Error while running startup `after` hooks...")
	}

	if err := http.ListenAndServe(net.JoinHostPort("", port), nil); err != nil {
		log.Fatalln("Could not startup server...", err)
	}
}
