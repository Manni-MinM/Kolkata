package main

import (
    "os"
    "log"

    "kolkata/pkg/quote"
    "kolkata/pkg/database"
	"kolkata/internal/server"
	"kolkata/internal/middleware"
)

func main() {
    db, err := database.NewRedis()
    if err != nil {
        log.Fatal(err)
    }

    logger := middleware.NewLogger(log.New(os.Stdout, "", log.LstdFlags))

    service := quote.NewHTTP(db)
    newHandler := logger(quote.NewQuoteHandler(service))
    searchHandler := logger(quote.SearchQuoteHandler(service))

    router := server.NewRouter()
    router.AddRoute("/new-quote", newHandler.ServeHTTP)
    router.AddRoute("/search-quote", searchHandler.ServeHTTP)

    err = server.ListenAndServe(":8000", router)
    if err != nil {
        log.Fatal(err)
    }
}

