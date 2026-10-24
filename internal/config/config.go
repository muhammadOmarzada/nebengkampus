package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config adalah struct (wadah data) yang menampung seluruh konfigurasi aplikasi
type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// Load bertugas membaca file .env dan mengembalikan pointer ke struct Config
func Load() (*Config, error) {
	// godotenv.Load() akan mencoba membaca file .env jika ada.
	// Jika file .env tidak ditemukan (misalnya saat dideploy ke cloud server),
	// kita tidak ingin program error, melainkan tetap lanjut membaca environment bawaan server.
	_ = godotenv.Load()

	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD"), // Password wajib dari env
		DBName:     getEnv("DB_NAME", "nebengkampus"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Validasi dasar: pastikan password database tidak kosong
	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD wajib diisi di file .env")
	}

	return cfg, nil
}

// ConnectionString menyusun format URL koneksi yang dipahami oleh PostgreSQL dan pgx
// Contoh format: postgres://postgres:password@localhost:5432/nebengkampus?sslmode=disable
func (c *Config) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

// getEnv adalah fungsi pembantu: membaca variabel env, jika kosong pakai nilai fallback/default
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
