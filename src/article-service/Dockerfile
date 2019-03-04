FROM golang:1.10

WORKDIR /go/src/article-service

COPY . .

RUN go install -v ./...

CMD ["article-service"]



