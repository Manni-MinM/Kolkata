package middleware

import (
    "io"
    "log"
    "bufio"
    "bytes"
    "errors"
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestLogger(t *testing.T) {
    reader, writer := io.Pipe()
    logger := NewLogger(log.New(writer, "", 0))

    panicHandler := http.HandlerFunc(func(http.ResponseWriter, *http.Request){
        err := errors.New("some-random-error")
        defer panic(err)
    })
    handler := logger(panicHandler)

    go func() {
        handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))
        writer.Close()
    }()

    buf := &bytes.Buffer{}
    sc := bufio.NewScanner(reader)
    for sc.Scan() {
        buf.WriteString(sc.Text())
    }

    result := buf.String()
    if expected := "some-random-error"; result != expected {
        t.Fatalf("expected result to be %v but got %v", expected, result)
    }
}

