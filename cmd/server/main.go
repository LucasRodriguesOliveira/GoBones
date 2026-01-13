package main

import (
	GoBones "github.com/LucasRodriguesOliveira/GoBones/internal/http"

	"fmt"
	"log"
	"net/http"
)

func Handler(r *GoBones.Request) *GoBones.Response {
	Headers := map[string]string{
		"Content-Type": "text/plain",
	}

	return &GoBones.Response{
		Status:  http.StatusOK,
		Headers: Headers,
		Body:    append([]byte("Hello, "), r.Body...),
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req, err := GoBones.ParseRequest(r)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

    res := GoBones.Logger(Handler)(req)

		GoBones.WriteResponse(w, res)
	})

	fmt.Println("=> Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
