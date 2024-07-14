FROM golang:1.22-alpine

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server


CMD ["./main"]
