package quote

import (
    "net/http"
)

func NewQuote(w http.ResponseWriter, req *http.Request) {
    q := quote{}
    quote, err := q.newQuote()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write([]byte(quote))
}

func ExistingQuote(w http.ResponseWriter, req *http.Request) {
    // TODO

    return
}

