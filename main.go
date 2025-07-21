package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// GORM bağlantısını başlat
	gormDB = connectDB()

	// Sahte veri ekleme: Her 5 saniyede bir
	go func() {
		for {
			val := 20 + rand.Float64()*10 // 20-30 arası sahte sıcaklık değeri
			gormDB.Create(&SensorData{Value: val})
			time.Sleep(5 * time.Second)
		}
	}()

	// Endpointler
	http.HandleFunc("/data/latest", getLatestData)          // son kaydı JSON olarak döndürür
	http.HandleFunc("/data/all", getAllData)                // son 50 sensör kaydını alır
	http.Handle("/", http.FileServer(http.Dir("./static"))) // static klasöründeki HTML dosyalarını sunar

	log.Println("Server 8080 portunda çalışıyor...")
	http.ListenAndServe(":8080", nil)
}

// En son sensör verisini döndürür
func getLatestData(w http.ResponseWriter, r *http.Request) {
	var d SensorData
	// created_at'e göre en yeni veriyi çek
	result := gormDB.Order("created_at desc").First(&d) // desc ile sıraladığı verilerden ilkini alır
	if result.Error != nil {
		http.Error(w, "Veri bulunamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(d)
}

// Son 50 veriyi döndürür
func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data []SensorData
	// created_at'e göre ters sırada, 50 kayıt
	result := gormDB.Order("created_at desc").Limit(50).Find(&data) // sıraladığı datalardanson 50 sini alır
	if result.Error != nil {
		http.Error(w, "Veriler alınamadı", 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
