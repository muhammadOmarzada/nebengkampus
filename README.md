# NebengKampus

**Aplikasi Ride-Sharing & Cost-Sharing Mahasiswa Berbasis Go dan PostGIS**

NebengKampus adalah project backend yang bertujuan membantu mahasiswa menemukan tebengan menuju kampus dan berbagi biaya perjalanan secara transparan. Project ini dikembangkan sebagai sarana belajar backend engineering menggunakan Go, PostgreSQL, dan PostGIS.

> **Project Status:** Planning / In Development

## Tujuan Project

NebengKampus bertujuan untuk:

- Menghubungkan mahasiswa yang membutuhkan transportasi dengan mahasiswa yang memiliki kendaraan dan kursi kosong.
- Mempermudah pencarian perjalanan berdasarkan lokasi, tujuan kampus, dan jadwal keberangkatan.
- Menyediakan perhitungan biaya bensin dan parkir yang transparan untuk perjalanan bersama.
- Membangun backend yang aman, terstruktur, dan dapat dikembangkan menggunakan teknologi yang relevan dengan industri.

Project ini berfokus pada ride-sharing antarmahasiswa, bukan layanan transportasi komersial seperti ojek online.

## Problem Statement

Mahasiswa yang tidak memiliki kendaraan pribadi sering mengalami kesulitan mendapatkan transportasi menuju kampus, terutama ketika jadwal keberangkatan dan lokasi tempat tinggal mereka berbeda-beda. Di sisi lain, mahasiswa yang membawa kendaraan pribadi terkadang memiliki kursi kosong saat berangkat ke kampus. NebengKampus dikembangkan sebagai aplikasi ride-sharing dan cost-sharing yang menghubungkan mahasiswa dengan tujuan perjalanan yang searah, sehingga mereka dapat berbagi kendaraan sekaligus membagi biaya bensin dan parkir secara transparan dan adil.

## MVP Features

Versi awal NebengKampus berfokus pada tiga fitur utama:

### 1. Trip Management

Driver dapat:

- Membuat perjalanan menuju kampus.
- Menentukan lokasi asal, tujuan, dan jadwal keberangkatan.
- Menentukan kapasitas kursi yang tersedia.
- Memasukkan estimasi biaya bensin dan parkir.
- Melihat dan mengelola perjalanan yang dibuat.

### 2. Geospatial Trip Search

Penumpang dapat:

- Mencari perjalanan menuju kampus yang dituju.
- Menemukan perjalanan berdasarkan lokasi penjemputan.
- Menyaring perjalanan berdasarkan jadwal keberangkatan.
- Melihat perjalanan yang lokasinya relevan menggunakan PostgreSQL PostGIS.

### 3. Booking & Cost Sharing

Penumpang dapat:

- Memesan kursi pada perjalanan yang tersedia.
- Melihat estimasi biaya patungan.
- Memeriksa status booking.

Backend akan memvalidasi ketersediaan kursi dan menghitung pembagian biaya berdasarkan aturan yang ditentukan.

## Main Database Tables

Rancangan awal database menggunakan enam tabel utama:

| Table | Description |
|---|---|
| `users` | Menyimpan profil mahasiswa, email kampus, dan password hash. |
| `trips` | Menyimpan perjalanan, driver, lokasi asal dan tujuan, jadwal, kapasitas, serta estimasi biaya. |
| `bookings` | Menyimpan pemesanan kursi oleh penumpang dan status booking. |
| `payments` | Menyimpan informasi transaksi pembayaran dan statusnya. |
| `wallets` | Menyimpan saldo internal driver untuk pengembangan fitur dompet. |
| `ledger_entries` | Menyimpan riwayat perubahan saldo dan referensi transaksi. |

> Catatan: tabel pembayaran dan dompet disiapkan untuk pengembangan berikutnya. Implementasi pembayaran dan pencairan dana belum termasuk dalam MVP awal.

## Tech Stack

| Technology | Purpose |
|---|---|
| Go (Golang) | Backend REST API dan business logic |
| PostgreSQL | Relational database |
| PostGIS | Geospatial query dan pencarian perjalanan berdasarkan lokasi |
| Docker | Menjalankan dan mengelola lingkungan development |
| DBeaver | Database management dan SQL exploration |
| Postman | API testing |
| Midtrans Sandbox | Rencana integrasi simulasi pembayaran pada fase pengembangan lanjutan |
| Git & GitHub | Version control dan dokumentasi project |

## Development Roadmap

- [ ] Perencanaan project dan database
- [ ] Desain ERD dan PostgreSQL schema
- [ ] Setup Go backend dan Docker
- [ ] Implementasi registrasi dan autentikasi mahasiswa
- [ ] Implementasi Trip Management API
- [ ] Implementasi pencarian perjalanan dengan PostGIS
- [ ] Implementasi booking dan validasi kapasitas kursi
- [ ] Implementasi perhitungan cost-sharing
- [ ] Integrasi Midtrans Sandbox dan internal wallet ledger
- [ ] Automated testing dan dokumentasi API

## Learning Goals

Melalui project ini, saya ingin mempelajari dan menerapkan:

- REST API development menggunakan Go.
- Relational database design dan SQL.
- Geospatial database menggunakan PostGIS.
- Transaction handling dan concurrency safety.
- Payment webhook dan idempotent transaction processing.
- Backend testing, dokumentasi, dan software engineering best practices.

## License

Project ini dibuat untuk tujuan pembelajaran dan portofolio. Informasi lisensi akan ditentukan sebelum project didistribusikan secara publik.