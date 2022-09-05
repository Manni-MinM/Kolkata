package quote

import (
	"io"
	"net/http"

	"github.com/buger/jsonparser"
)

type quoteService interface {
    newQuote() (string, error)
    existingQuote() (string, error)
}

type quote struct{}

func (quote) newQuote() (string, error) {
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

func (quote) existingQuote() (string, error) {
    // TODO

    return "", nil
}

