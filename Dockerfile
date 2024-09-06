FROM golang:1.23.1-bookworm
WORKDIR /app
COPY . /app
RUN go build
ENTRYPOINT ["./suninfo-notification"]