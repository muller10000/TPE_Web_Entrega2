# Etapa 1: build
FROM golang:1.25 AS builder
WORKDIR /app

# Copiar módulos y dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario
RUN apt-get update && apt-get install -y curl && \
rm -rf /var/lib/apt/lists/* && go build -o peliculas-api .

ENV PORT=8080
EXPOSE 8080

CMD ["/app/peliculas-api"]