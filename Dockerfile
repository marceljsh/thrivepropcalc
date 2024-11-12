FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/thrivepropcalc ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /bin/thrivepropcalc .

COPY input.txt ./
COPY public/ ./public/

CMD ["./thrivepropcalc"]