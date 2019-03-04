#FROM microhq/micro
#CMD ["api", "--handler=rpc" ,  "--address=:8080" ,  "--namespace=gomicrodemo1"  ]

FROM golang:1.10

WORKDIR /go/src/micro

COPY . .

RUN go install -v ./...

#CMD ["micro"]
ENTRYPOINT ["micro"]
