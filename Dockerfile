FROM golang:latest

WORKDIR /app

COPY cmd/main.go .

RUN go build -o /app/bin/propcalc main.go

ENTRYPOINT ["/app/bin/propcalc"]