package main

import (
	"fmt"

	"github.com/LucasRodriguesOliveira/GoBones/core/server"
	"github.com/LucasRodriguesOliveira/GoBones/core/server/pipeline"
	GoBones "github.com/LucasRodriguesOliveira/GoBones/internal/http"

	"net/http"
)

func Handler(req *GoBones.Request, res *GoBones.Response) error {
	res.Body = append([]byte("Hello, "), req.Body...)
	res.Headers = map[string]string{"Content-Type": "text/plain"}
	res.Status = http.StatusOK

	return nil
}

func main() {
	server := &server.Server{
		Config:  server.Config{Port: 8080},
		Startup: pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]{},
		Hooks:   pipeline.Pipeline[func(req *GoBones.Request, res *GoBones.Response) error]{},
	}

	server.Hooks.Register(GoBones.Logger, pipeline.PIPELINE_REGISTER_BEFORE)

	fmt.Println("Starting server...")
	server.Start()
}
