package main

import (
    "log"

    "kolkata/pkg/quote"
	"kolkata/internal/server"
)

func main() {
    r := server.NewRouter()
    r.AddRoute("/new-quote", quote.NewQuote)
    r.AddRoute("/existing-quote", quote.ExistingQuote)

    err := server.ListenAndServe(":8000", r)
    if err != nil {
        log.Fatal(err)
    }
}

