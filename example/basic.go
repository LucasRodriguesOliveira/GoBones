package main

import (
	"fmt"
	"net/http"

	"github.com/LucasRodriguesOliveira/GoBones/core/pipeline"
	GB "github.com/LucasRodriguesOliveira/GoBones/core/server"
	H "github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

func Hello(req *H.Request, res *H.Response) error {
	res.Ok(H.J{"message": "Hello World!"})

	return nil
}

func main() {
	app := GB.New(8080)

	app.Hooks.Register(H.Logger, pipeline.PIPELINE_REGISTER_BEFORE)
	app.Router.Register("/", http.MethodGet, Hello)

	fmt.Println("Starting server at port 8080...")
	app.Start()
}
