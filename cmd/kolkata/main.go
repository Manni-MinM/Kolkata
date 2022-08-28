package main

import (
    "log"

	"kolkata/internal/server"
    "kolkata/internal/service/quote"
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

