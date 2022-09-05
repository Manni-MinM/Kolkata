package main

import (
    "log"

    "kolkata/pkg/quote"
    "kolkata/pkg/database"
	"kolkata/internal/server"
)

func main() {
    db, err := database.NewRedis()
    if err != nil {
        log.Fatal(err)
    }

    q := quote.NewHTTP(db)
    newHandler := quote.NewQuoteHandler(q)
    searchHandler := quote.SearchQuoteHandler(q)

    r := server.NewRouter()
    r.AddRoute("/new-quote", newHandler.ServeHTTP)
    r.AddRoute("/search-quote", searchHandler.ServeHTTP)

    err = server.ListenAndServe(":8000", r)
    if err != nil {
        log.Fatal(err)
    }
}

