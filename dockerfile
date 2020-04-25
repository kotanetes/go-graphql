FROM golang:alpine as builder
ADD . /go-graphql
WORKDIR /go-graphql 
RUN go get -t -v ./...
RUN go build -o app;

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go-graphql/app .
EXPOSE 9090
#RUN chmod -r 777
CMD ["./app"]