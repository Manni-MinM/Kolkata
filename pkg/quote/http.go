package quote

import (
	"io"
	"net/http"

	"github.com/buger/jsonparser"
)

type httpQuote struct{}

func NewHTTP() (QuoteService, error) {
    return &httpQuote{}, nil
}

func (q *httpQuote) GetQuote() (string, error) {
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

