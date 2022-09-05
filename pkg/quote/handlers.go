package quote

import (
    "net/http"
    "encoding/json"
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

        var id string
        err := json.NewDecoder(req.Body).Decode(&id)
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

