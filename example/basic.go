package main

import (
	"fmt"

	GB "github.com/LucasRodriguesOliveira/GoBones/core/server"
	"github.com/LucasRodriguesOliveira/GoBones/core/pipeline"
	"github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

func Hello(req *http.Request, res *http.Response) error {
	res.Ok(http.J{ "message": "Hello World!" })

	return nil
}

func main() {
	app := GB.New(8080)

	app.Hooks.Register(http.Logger, pipeline.PIPELINE_REGISTER_BEFORE)
	app.Router.Register("/", "GET", Hello)

	fmt.Println("Starting server at port 8080...")
	app.Start()
}
