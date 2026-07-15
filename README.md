# ⚡ TechMart — Full-Stack E-Commerce Application

TechMart adalah aplikasi e-commerce full-stack untuk penjualan produk elektronik, dibangun dari nol menggunakan **Golang** (backend), **Angular** (frontend), dan **MySQL** (database). Project ini mencakup alur bisnis e-commerce yang lengkap — mulai dari autentikasi, manajemen produk, keranjang belanja, checkout, pembayaran manual, hingga dashboard admin dengan statistik penjualan.

## 🔗 Live Demo

- **Frontend:** _(akan ditambahkan setelah deployment)_
- **Backend API:** _(akan ditambahkan setelah deployment)_


## 🛠️ Tech Stack

**Backend**
- Golang + [Gin](https://gin-gonic.com/) (web framework)
- [GORM](https://gorm.io/) (ORM)
- MySQL (database)
- JWT (autentikasi)
- Bcrypt (hashing password)

**Frontend**
- Angular 18 (standalone components)
- Angular Material concepts + custom CSS
- Chart.js (visualisasi data)
- RxJS

**Lainnya**
- RESTful API
- Role-Based Access Control (Admin & Customer)
- File upload (gambar produk, foto profil, bukti transfer)

## ✨ Fitur Utama

### Autentikasi & Keamanan
- Register & Login dengan JWT
- Role-based access control (Admin vs Customer)
- Route guard di frontend (proteksi halaman sesuai role)
- Lupa password dengan security question
- Ganti password dari halaman profil

### Untuk Customer
- Jelajahi katalog produk dengan search, filter kategori, dan pagination
- Halaman detail produk
- Keranjang belanja (tambah, ubah jumlah, hapus)
- Checkout & pembayaran manual (upload bukti transfer)
- Riwayat pesanan dengan tracking status
- Edit profil (nama, jenis kelamin, foto profil)
- Notifikasi in-app real-time (status pesanan, dll)

### Untuk Admin
- Dashboard dengan statistik toko (grafik pendapatan, produk terlaris, status pesanan)
- Kelola produk (CRUD lengkap + upload gambar)
- Kelola kategori produk
- Kelola pesanan (verifikasi pembayaran, update status)
- Kelola customer (lihat profil & riwayat belanja setiap customer)
- Admin tidak dapat berbelanja (role separation yang konsisten)

## 🗂️ Struktur Project

```
techmart-ecommerce/
├── backend/              # Golang REST API
│   ├── config/           # Konfigurasi database
│   ├── controllers/      # Business logic
│   ├── middleware/       # Auth & role middleware
│   ├── models/           # Struktur data (GORM models)
│   ├── routes/           # Definisi routing API
│   ├── utils/            # Helper functions (JWT, hashing, dll)
│   └── main.go
└── frontend/             # Angular application
    └── src/app/
        ├── auth/          # Login, Register, Forgot Password
        ├── dashboard/     # Dashboard utama
        ├── products/      # Katalog, Detail, Kelola Produk
        ├── cart/          # Keranjang belanja
        ├── orders/        # Riwayat & Kelola Pesanan
        ├── profile/       # Edit profil
        ├── users/         # Kelola Customer (admin)
        ├── stats/         # Dashboard Statistik (admin)
        └── notifications/ # Notifikasi in-app
```

## 🚀 Cara Menjalankan di Lokal

### Prasyarat
- [Go](https://go.dev/dl/) 1.22+
- [Node.js](https://nodejs.org/) 18+ & Angular CLI
- MySQL

### 1. Clone Repository
```bash
git clone https://github.com/Bronsonsiallagan/techmart-ecommerce.git
cd techmart-ecommerce
```

### 2. Setup Backend
```bash
cd backend
go mod tidy
```

Buat file `.env` (contoh ada di `.env.example`):
```env
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=techmart
JWT_SECRET=ganti_dengan_string_rahasia_anda
```

Buat database MySQL:
```sql
CREATE DATABASE techmart;
```

Jalankan server:
```bash
go run main.go
```
Backend akan berjalan di `http://localhost:8080`

### 3. Setup Frontend
Buka terminal baru:
```bash
cd frontend
npm install
ng serve
```
Frontend akan berjalan di `http://localhost:4200`

## 📋 API Endpoints (Ringkasan)

| Method | Endpoint | Deskripsi | Akses |
|--------|----------|-----------|-------|
| POST | `/api/auth/register` | Registrasi akun baru | Publik |
| POST | `/api/auth/login` | Login | Publik |
| GET | `/api/products` | Lihat semua produk | Publik |
| POST | `/api/products` | Tambah produk | Admin |
| GET | `/api/cart` | Lihat keranjang | Customer |
| POST | `/api/orders` | Checkout | Customer |
| GET | `/api/admin/orders` | Lihat semua pesanan | Admin |
| GET | `/api/admin/stats` | Statistik toko | Admin |

_Dokumentasi API lengkap dapat dilihat di kode `routes/routes.go`._

## 🎯 Highlight Teknis

- **Role-Based Access Control** diterapkan konsisten di backend (middleware) dan frontend (route guard), bukan hanya menyembunyikan UI
- **Data integrity protection** — produk yang sudah pernah dipesan tidak bisa dihapus untuk menjaga riwayat transaksi tetap utuh
- **Pagination** diterapkan di backend untuk performa yang lebih baik pada dataset besar
- **File upload** untuk gambar produk, foto profil, dan bukti transfer pembayaran
- **In-app notification system** dengan polling untuk update real-time-ish

## 👤 Kontak

**Bronson Siallagan**
GitHub: [@Bronsonsiallagan](https://github.com/Bronsonsiallagan)

---

_Project ini dibuat sebagai bagian dari portofolio pengembangan web full-stack._