FROM golang:alpine AS builder

WORKDIR /app
COPY ./ ./

WORKDIR /app/cmd/kolkata
RUN go build -o kolkata

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/kolkata/kolkata .

EXPOSE 8000
ENTRYPOINT ["./kolkata"]

