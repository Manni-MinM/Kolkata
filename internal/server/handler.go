package server

import (
    "net/http"

    "kolkata/internal/service"
)

func newQuote(w http.ResponseWriter, req *http.Request) {
    quoteService := service.Quote{}
    quote, err := quoteService.NewQuote()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write([]byte(quote))
}

func existingQuote(w http.ResponseWriter, req *http.Request) {
    // TODO

    return
}

