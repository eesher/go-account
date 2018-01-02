FROM alpine:latest

COPY ./go-account /root
WORKDIR /root

EXPOSE 8080
CMD  ["./go-account"]
