# Etapa de build
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Imagem final mínima
FROM gcr.io/distroless/base-debian13

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
