package server

import (
    "fmt"
	"net/http"

    "kolkata/internal/config"
)

type Handler func(http.ResponseWriter, *http.Request)

func ListenAndServe(conf config.ServerConfig, r Router) error {
    for _, route := range r.Routes {
        r.Mux.Handle(route.Path, http.HandlerFunc(route.Handler))
    }

    portString := fmt.Sprintf(":%v", conf.Port)
    return http.ListenAndServe(portString, r.Mux)
}

