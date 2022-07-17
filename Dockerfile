FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY cmd ./cmd


RUN go build cmd/main/main.go

EXPOSE 8080

CMD ["./main"]