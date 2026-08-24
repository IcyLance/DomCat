# Build
FROM golang:1.22-bookworm AS builder

WORKDIR /src

# Download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /domCat .

# Runtime
FROM debian:bookworm-slim

WORKDIR /app

# Install CA certificates for HTTPS API calls
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /domCat /usr/local/bin/domCat

ENTRYPOINT ["/usr/local/bin/domCat"]
