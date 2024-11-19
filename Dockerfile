FROM golang:alpine3.20 AS builder

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .
RUN go build -o /app ./cmd/*

FROM alpine:3.20

COPY --from=builder /app /app
COPY --from=builder /build/configs /configs
CMD ["/app"]
