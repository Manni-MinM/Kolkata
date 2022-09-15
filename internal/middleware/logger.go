package middleware

import (
    "log"
    "net/http"
)

func NewLogger(l *log.Logger) middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            defer func() {
                if r := recover(); r != nil {
                    l.Println(r)
                }
            }()

            next.ServeHTTP(w, req)
        })
    }
}

