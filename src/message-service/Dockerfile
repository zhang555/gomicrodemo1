FROM golang:1.10

WORKDIR /go/src/message-service

COPY . .

RUN go install -v ./...

CMD ["message-service"]

