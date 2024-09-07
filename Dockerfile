FROM golang:1.23-alpine AS build
ENV CGO_ENABLED=1
RUN apk add gcc gcc-go libc-dev
WORKDIR /app
COPY . /app
RUN go build

FROM alpine:3
WORKDIR /app
COPY --from=build /app/suninfo-notification /app/suninfo-notification
ENTRYPOINT ["/app/suninfo-notification"]