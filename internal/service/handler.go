package service

import (
	"io"
	"net/http"

	"github.com/buger/jsonparser"
)

type Quote struct{}

func (Quote) NewQuote() (string, error) {
    resp, err := http.Get("https://api.quotable.io/random")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return jsonparser.GetString(body, "content")
}

func (Quote) ExistingQuote() (string, error) {
    // TODO

    return "", nil
}

