FROM golang:1.23-alpine
ENV CGO_ENABLED=1
RUN apk add gcc-go libc-dev
WORKDIR /app
COPY . /app
RUN go build
ENTRYPOINT ["./suninfo-notification"]