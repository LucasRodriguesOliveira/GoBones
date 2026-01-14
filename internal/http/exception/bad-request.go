package exception

import "net/http"

func BadRequestException(w http.ResponseWriter) {
  http.Error(w, "Bad Request", http.StatusBadRequest)
}
