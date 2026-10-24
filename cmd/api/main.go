package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadOmarzada/nebengkampus/internal/config"
	"github.com/muhammadOmarzada/nebengkampus/internal/database"
)

// Application menampung dependensi yang dibutuhkan oleh handler HTTP
type Application struct {
	config *config.Config
	db     *pgxpool.Pool
}

func main() {
	// 1. Baca konfigurasi dari .env
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}
	log.Println("Konfigurasi berhasil dimuat.")

	// 2. Hubungkan ke database PostgreSQL
	dbPool, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Gagal menghubungkan database: %v", err)
	}
	defer dbPool.Close() // Pastikan koneksi database ditutup saat aplikasi dimatikan

	log.Println("Berhasil terhubung ke database PostgreSQL!")

	// 3. Inisialisasi struct aplikasi
	app := &Application{
		config: cfg,
		db:     dbPool,
	}

	// 4. Siapkan Router HTTP bawaan Go
	mux := http.NewServeMux()

	// Daftarkan route GET /health
	mux.HandleFunc("GET /health", app.healthCheckHandler)

	// 5. Konfigurasi dan nyalakan HTTP Server
	serverAddr := ":" + cfg.ServerPort
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server NebengKampus berjalan di http://localhost%s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server berhenti karena error: %v", err)
	}
}

// healthCheckHandler merespons status server dan koneksi database dalam bentuk JSON
func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Lakukan ping cepat ke database (timeout 2 detik)
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "connected"
	if err := app.db.Ping(ctx); err != nil {
		dbStatus = "unreachable"
	}

	// Buat payload respons JSON
	response := map[string]string{
		"status":   "available",
		"database": dbStatus,
		"system":   "NebengKampus API",
	}

	// Kirim header dan data JSON ke klien/browser
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}