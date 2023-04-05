FROM golang:1.8 as builder
COPY / /app
WORKDIR /app
RUN GO111MODULE=on CGO_ENABLED=0 GOSUMDB=off GOOS=linux go build -a -o /go/bin/quck .



FROM alpine:3.15
EXPOSE 8888
COPY --from=builder /go/bin/quck ./
CMD ["./quck","web"]