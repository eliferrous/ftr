FROM golang:1.23-alpine AS builder

WORKDIR /src
ENV CGO_ENABLED=0

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /ftr ./cmd/ftr

FROM alpine:3.19

RUN apk add --no-cache mtr curl ca-certificates libc6-compat

COPY --from=builder /ftr /app/ftr
RUN chmod +x /app/ftr


COPY --from=flyio/flyctl flyctl /usr/bin

EXPOSE 8080
CMD ["flyctl", "mcp", "wrap", "--mcp", "/app/ftr", "--port", "8080"]
