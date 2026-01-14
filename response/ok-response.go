package response

import (
	"fmt"
	"net/http"

	GoBones "github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

func OkResponse(res *GoBones.Response, message string) {
  res.Headers["Content-Type"] = "application/json"
  res.Status = http.StatusOK
  res.Body = fmt.Appendf(nil, "{ %q: %q }", "message", message)
}
