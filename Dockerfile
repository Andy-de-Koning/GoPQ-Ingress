# --- Stage 1: Builder ---
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gopq-ingress main.go

# --- Stage 2: Runner ---
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/gopq-ingress .
RUN mkdir -p /root/.local/share/certmagic
EXPOSE 80 443
CMD ["./gopq-ingress"]
