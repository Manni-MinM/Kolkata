package quote

type QuoteService interface {
    NewQuote() (string, error)
    SearchQuote(string) (string, error)
}

