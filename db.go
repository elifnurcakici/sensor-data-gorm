package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// PostgreSQL'e GORM ile bağlanır
func connectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=843038 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	// Tabloyu otomatik oluştur (eğer yoksa)
	db.AutoMigrate(&SensorData{})

	return db
}
