package quote

import (
    "io"
    "bytes"
    "testing"
    "net/http"
    "net/http/httptest"

    "kolkata/pkg/database"
)

func TestNewQuoteHandler(t *testing.T) {
    db := database.NewMock()

    httpQS := NewHTTP(db)
    handler := NewQuoteHandler(httpQS)

    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/quotes/new", nil)
    handler.ServeHTTP(w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected status code to be 200 but got %v", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        t.Fatal("couldn't parse response body")
    } else if body == nil {
        t.Fatal("expected response body but got nil")
    }
}

func TestSearchQuoteHandler(t *testing.T) {
    db := database.NewMock()
    db.Set("some-random-key", "some-random-value")

    httpQS := NewHTTP(db)
    handler := SearchQuoteHandler(httpQS)

    w := httptest.NewRecorder()
    reqBody := []byte(`{"id": "some-random-key"`)
    req := httptest.NewRequest(http.MethodPost, "/quotes/search", bytes.NewBuffer(reqBody))
    handler.ServeHTTP(w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected status code to be 200 but got %v", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        t.Fatal("couldn't parse response body")
    }

    result := string(body)
    if expected := "some-random-value"; result != expected {
        t.Fatalf("expected response body to be %v but got %v", expected, result)
    }
}

