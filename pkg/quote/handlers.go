package quote

import (
    "io"
    "net/http"

    "github.com/buger/jsonparser"
)

func NewQuoteHandler(q QuoteService) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        if req.Method != http.MethodGet {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        id, err := q.NewQuote()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(id))
    })
}

func SearchQuoteHandler(q QuoteService) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        if req.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        body, err := io.ReadAll(req.Body)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        id, err := jsonparser.GetString(body, "id")
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        text, err := q.SearchQuote(id)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(text))
    })
}

