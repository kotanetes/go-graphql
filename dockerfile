FROM golang:alpine
ADD . /go-graphql
WORKDIR /go-graphql 
RUN go get -t -v ./...
RUN go build -o app;
CMD ["./app"]