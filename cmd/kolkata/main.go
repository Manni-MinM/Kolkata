package main

import (
    "log"

	"kolkata/internal/server"
)

func main() {
    err := server.ListenAndServe(":8000")
    if err != nil {
        log.Fatal(err)
    }
}

