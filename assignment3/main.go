package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Struct untuk menyimpan status air dan angin
type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Jalankan goroutine untuk mengupdate file setiap 15 detik
	go updateFile()

	// Handle HTTP route untuk menyajikan halaman HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handle HTTP route untuk menyajikan file status JSON
	http.HandleFunc("/status.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("status.json")
		if err != nil {
			http.Error(w, "Error reading status file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Content-Type", "text/event-stream")

		w.Write(data)
	})

	// Mulai server HTTP
	http.ListenAndServe(":8080", nil)
}

func updateFile() {
	for {
		// Generate angka random antara 1-100 untuk air dan angin
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Buat struct Status baru
		status := Status{
			Water: water,
			Wind:  wind,
		}

		// Marshal struct menjadi JSON
		jsonData, err := json.MarshalIndent(status, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		// Tulis data JSON ke file
		err = os.WriteFile("status.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}

		// Tunggu 15 detik sebelum mengupdate kembali
		time.Sleep(15 * time.Second)
	}
}
