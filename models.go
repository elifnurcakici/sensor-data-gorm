package main

import "time"

// SensorData: Veritabanındaki sensör verilerini temsil eden model
type SensorData struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}
