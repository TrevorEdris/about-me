FROM golang:1.17-alpine AS builder

RUN apk update && apk add musl-dev gcc build-base ca-certificates

WORKDIR /app

COPY . .
RUN go build -ldflags "-linkmode external -extldflags \"-static\" -s -w $LDFLAGS" -o api cmd/api/main.go

# Copy the binary from "builder" into a scratch container to reduce the overall size of the image
FROM scratch AS final

ENTRYPOINT ["/app/api"]
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/api /app/api
