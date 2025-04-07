# build stage
FROM golang:1.24-alpine as builder

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -o main

# runtime stage
FROM debian:12-slim

RUN apt-get update -y \
    && apt-get install -y --no-install-recommends openssl ca-certificates \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /build/main /app/main

ENTRYPOINT ["./main"]