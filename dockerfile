FROM golang:1.18 as builder
COPY / /app
WORKDIR /app
ENV GO111MODULE=on GOPROXY="https://goproxy.cn,direct"  CGO_ENABLED=0 GOSUMDB=off GOOS=linux
RUN go mod tidy
RUN go build -a -o /go/bin/quck .



FROM alpine:3.15
EXPOSE 8888
COPY --from=builder /go/bin/quck .
CMD ["/quck","web"]