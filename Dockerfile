# Builder stage
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o arithmancy

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/arithmancy .
COPY templates /app/templates
COPY static /app/static
EXPOSE 8080
CMD ["./arithmancy"] 