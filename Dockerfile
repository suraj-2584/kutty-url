FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5001
ENV HOST="kutty-url.in"
CMD ["./main"]