package service

type QuoteService interface {
    NewQuote() (string, error)
    ExistingQuote() (string, error)
}

