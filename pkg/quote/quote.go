package quote

type QuoteService interface {
    GetQuote() (string, error)
}

