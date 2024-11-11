FROM golang:latest

WORKDIR /app

COPY cmd/main.go .

RUN go build -o /app/bin/thrivepropcalc main.go

ENTRYPOINT ["/app/bin/thrivepropcalc"]