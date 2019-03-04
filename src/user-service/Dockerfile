FROM golang:1.10

WORKDIR /go/src/user-service

COPY . .

RUN go install -v ./...

CMD ["user-service"]
