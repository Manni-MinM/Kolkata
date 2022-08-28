package server

import (
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

func ListenAndServe(port string, r Router) error {
    for _, route := range r.Routes {
        r.Mux.Handle(route.Path, http.HandlerFunc(route.Handler))
    }

    return http.ListenAndServe(port, r.Mux)
}

