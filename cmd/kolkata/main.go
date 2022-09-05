package main

import (
	"log"
	"os"

	"kolkata/internal/config"
	"kolkata/internal/middleware"
	"kolkata/internal/server"
	"kolkata/pkg/database"
	"kolkata/pkg/quote"
)

func main() {
    conf, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    db, err := database.NewRedis(conf.Redis)
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

    err = server.ListenAndServe(conf.Server, router)
    if err != nil {
        log.Fatal(err)
    }
}

