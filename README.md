# GoBones

A simplistic HTTP web server down to bare**Bones**

## Usage

Just create a new server object and start it!

```go
func main() {
  app := server.New(8080)

  fmt.Println("Starting sever at port 8080...")
  app.Start()
}
```

```bash
> Starting server at port 8080...
```

simple as that

### Routes

Want to add a route? Just register it to the router object inside the server and pass a handler

```go
package main

import (
  "fmt"

  GB  "github.com/LucasRodriguesOliveira/GoBones/core/server"
  "github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

func Hello(req *http.Request, res *http.Response) error {
  res.Ok(http.J{ "message": "Hello" })

  return nil
}

func main() {
  app := GB.New(8080)
  app.Router.Register("/", "Get", Hello)

  fmt.Println("Starting server at port 8080...")
  app.Start()
}
```

### Pipelines

For middlewares, is just as simples as adding a route. Access the specific Pipeline that
best suit your situation. If you want to add a hook to every request pipeline, go for it!

```go
package main

import (
  "fmt"

  GB "github.com/LucasRodriguesOliveira/GoBones/core/server"
  "github.com/LucasRodriguesOliveira/GoBones/core/pipeline"
  "github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

func main() {
  app := GB.New(8080)

  app.Hooks.Register(http.Logger, pipeline.PIPELINE_REGISTER_BEFORE)

  fmt.Println("Starting server at port 8080...")
  app.Start()
}
```

That's it. GoBones is made to be simple.

Try it on!
