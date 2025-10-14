FROM golang:1.25-alpine

WORKDIR /app

# Copiar go.mod y go.sum primero para cache de dependencias
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar todo el proyecto
COPY . .

# Compilar el binario
RUN go build -o peliculas-api main.go

# Ejecutar el binario
CMD ["./peliculas-api"]