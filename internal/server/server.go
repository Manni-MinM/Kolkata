package server

import (
    "net/http"
)

func ListenAndServe(port string) error {
    mux := http.NewServeMux()

    newQuoteHandler := http.HandlerFunc(newQuote)
    existingQuoteHandler := http.HandlerFunc(existingQuote)

    mux.Handle("/new-quote", newQuoteHandler)
    mux.Handle("/existing-quote", existingQuoteHandler)

    return http.ListenAndServe(port, mux)
}

