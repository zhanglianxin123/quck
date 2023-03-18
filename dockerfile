FROM alpine:3.15
EXPOSE 8888
COPY ./quck ./
CMD ["./quck","web"]