package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhammadOmarzada/nebengkampus/internal/config"
)

// Connect menginisialisasi connection pool ke PostgreSQL menggunakan pgxpool
func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	// 1. Buat context dengan batas waktu (timeout) 5 detik
	// Jika dalam 5 detik database tidak merespons, proses koneksi akan dibatalkan otomatis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // cancel() dipanggil untuk membersihkan resource context setelah fungsi selesai

	// 2. Ambil connection string dari package config yang sudah kita buat sebelumnya
	connStr := cfg.ConnectionString()

	// 3. Konfigurasi pgxpool dari connection string
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("gagal memproses konfigurasi database: %w", err)
	}

	// Anda bisa menentukan batas jumlah koneksi pool (opsional tapi praktik yang baik)
	poolConfig.MaxConns = 10                     // Maksimal 10 koneksi terbuka sekaligus
	poolConfig.MinConns = 2                      // Minimal 2 koneksi selalu standby
	poolConfig.MaxConnIdleTime = 5 * time.Minute // Koneksi nganggur ditutup setelah 5 menit

	// 4. Inisialisasi pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat connection pool: %w", err)
	}

	// 5. Lakukan "Ping" nyata ke database untuk memverifikasi autentikasi & jaringan
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // Tutup pool jika ping gagal agar tidak ada memory leak
		return nil, fmt.Errorf("gagal terhubung (ping) ke database: %w", err)
	}

	return pool, nil
}