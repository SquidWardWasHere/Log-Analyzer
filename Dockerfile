# Build Asamasi (Latest kullanarak sürüm hatasını çözüyoruz)
FROM golang:alpine AS builder

WORKDIR /app

# Bağımlılıkları kopyala ve indir
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodları kopyala
COPY . .

# Binary dosyayı derle (CGO kapalı, static link)
RUN CGO_ENABLED=0 GOOS=linux go build -o log-tool cmd/main.go

# Final Asaması (Çalıştırma - Çok küçük boyutlu)
FROM alpine:latest

WORKDIR /root/

# Derlenmiş programı ve config'i al
COPY --from=builder /app/log-tool .
COPY --from=builder /app/config ./config

# Logların ve Raporun olacağı klasörleri oluştur
RUN mkdir /logs /output

# Programı çalıştır
CMD ["./log-tool"]

