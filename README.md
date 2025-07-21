# Sensor Data API (GORM Version)

Bu proje, sahte sensör verilerini (20-30 °C arası sıcaklık) her 5 saniyede bir PostgreSQL veritabanına kaydeder ve bu verileri bir REST API ve basit bir web arayüzü aracılığıyla görüntüler.

## Özellikler
- Go ile yazılmış backend
- GORM kullanılarak PostgreSQL ile bağlantı
- Her 5 saniyede bir otomatik sahte veri ekleme
- `/data/latest` endpoint’i ile en son değeri alma
- `/data/all` endpoint’i ile son 50 değeri listeleme
- Basit HTML frontend (`static/index.html`)

## Kurulum

1. Go ve PostgreSQL'in yüklü olduğundan emin olun.
2. PostgreSQL'de bir veritabanı oluşturun:
   ```bash
   createdb testdb
   
3. Projeyi klonlayın:
git clone https://github.com/elifnurcakici/sensor-data-gorm.git
cd sensor-data-gorm

4. Gerekli paketleri yükleyin:
go mod tidy

5. Projeyi çalıştırın:
go run .https://docs.github.com/github/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax

# Kullanım

Tarayıcıdan http://localhost:8080/ adresine gidin.
En son veri için: http://localhost:8080/data/latest
Son 50 veri için: http://localhost:8080/data/all

# Teknolojiler

Go
GORM (PostgreSQL ORM)
PostgreSQL
Basit HTML + JavaScript frontend
