package quote

import (
    "io"
    "net/http"

    "kolkata/pkg/database"

    "github.com/google/uuid"
    "github.com/buger/jsonparser"
)

type httpQuote struct{
    db    database.Database
}

func NewHTTP(db database.Database) QuoteService {
    return &httpQuote{db}
}

func (q *httpQuote) NewQuote() (string, error) {
    resp, err := http.Get("https://api.quotable.io/random")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    text, err := jsonparser.GetString(body, "content")
    if err != nil {
        return "", err
    }

    id := uuid.NewString()
    return q.db.Set(id, text)
}

func (q *httpQuote) SearchQuote(id string) (string, error) {
    return q.db.Get(id)
}

