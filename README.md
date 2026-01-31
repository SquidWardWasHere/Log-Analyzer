# CLI Tabanlı Log Analiz ve Uyarı Aracı

Bu proje, Go (Golang) kullanılarak geliştirilmiş, kural tabanlı bir log analiz sistemidir. Docker üzerinde çalışır ve host sistemdeki logları gerçek zamanlı izleyerek (tailing) veya geçmişe dönük tarayarak CSV raporu oluşturur.

**Hazırlayan:** Hasan Ali Kahraman

## Özellikler
* **Modüler Yapı:** Clean Architecture prensiplerine uygun klasörleme.
* **Volume Mapping:** Host makinedeki logları container içine bağlar.
* **Encoding Desteği:** Windows loglarında oluşan karakter bozukluklarını (UTF-16/Null Byte) otomatik temizler.
* **Canlı İzleme:** `tail -f` benzeri yapıyla anlık uyarı üretir.

## Kurulum ve Çalıştırma

### 1. Docker Imajını Oluşturma
Proje dizininde terminali açın:

```bash
docker build -t log-analyzer .

## Windows İçin Kurulum ve Çalıştırma
Docker Imajını Oluşturma

```bash
docker build -t log-analyzer .

Konteyner başlatma

```bash
docker run -it --rm -v //c/Logs:/logs:ro -v "${PWD}/output:/output" log-analyzer

## Linux İçin Kurulum ve Çalıştırma
Docker Imajını Oluşturma

```bash
docker build -t log-analyzer .

Konteyner Başlatma

```bash
docker run -it --rm -v "$(pwd)/logs:/logs:ro" -v "$(pwd)/output:/output" log-analyzer

Not: Linux üzerinde sistem loglarını okumak için Dockerı sudo ile çalıştırmamız gerekebilir.