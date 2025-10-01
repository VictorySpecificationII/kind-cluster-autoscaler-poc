# Build stage
FROM golang:1.25.1 as builder
WORKDIR /app
COPY autoscaler.go .
RUN go build -o autoscaler autoscaler.go

# Minimal runtime
FROM alpine:3.18
RUN apk add --no-cache bash curl wget tar

# Install kubectl manually
RUN wget -qO /usr/local/bin/kubectl https://dl.k8s.io/release/$(wget -qO- https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x /usr/local/bin/kubectl

COPY --from=builder /app/autoscaler /autoscaler

ENTRYPOINT ["/autoscaler"]
