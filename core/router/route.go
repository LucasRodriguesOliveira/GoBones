package router

import "github.com/LucasRodriguesOliveira/GoBones/internal/http"

type Route map[string]http.HandlerFunc
