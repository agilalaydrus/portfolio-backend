# Gunakan image golang
FROM golang:1.24

# Set workdir di dalam container
WORKDIR /app

# Copy go mod dan go sum dulu biar caching optimal
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o main .

# Expose port sesuai di .env (misal: 8090)
EXPOSE 8090

# Jalankan aplikasi
CMD ["./main"]
