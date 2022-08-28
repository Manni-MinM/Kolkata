package server

import (
    "net/http"
)

type Route struct {
    Path       string
    Handler    Handler
}

type Router struct {
    Mux       *http.ServeMux
    Routes    []Route
}

func NewRouter() Router {
    return Router{
        Mux: http.NewServeMux(),
        Routes: nil,
    }
}

func (r *Router) AddRoute(path string, handler Handler) {
    newRoute := Route{path, handler}
    r.Routes = append(r.Routes, newRoute)
}

