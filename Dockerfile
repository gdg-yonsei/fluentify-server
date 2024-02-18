FROM bufbuild/buf:latest AS buf
WORKDIR /build
COPY . .
RUN buf generate idl/proto

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.21-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
COPY --from=buf /build/gen ./gen
RUN go mod download
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/main /build/.env /app/
ENTRYPOINT ["./main"]
