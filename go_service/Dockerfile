FROM golang:1.21

RUN mkdir /go-service

WORKDIR /go-service

COPY go.mod go.sum ./

COPY .env .

RUN go mod download

COPY . .

CMD go run cmd/main.go