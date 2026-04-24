# 🧴 SkinCare Routine App

Aplikasi web fullstack untuk membantu pengguna menentukan rutinitas skincare harian berdasarkan kondisi kulit mereka.

## 🎯 Fitur Utama

- ✅ **Autentikasi Pengguna** - Register, Login, Logout dengan JWT
- ✅ **Analisis Kondisi Kulit** - Input jenis kulit, masalah, dan sensitivitas
- ✅ **Rekomendasi Skincare** - Rutinitas pagi dan malam yang dipersonalisasi
- ✅ **Penyimpanan Data** - Riwayat analisis dengan database
- ✅ **Dashboard** - Tampilan profil dan riwayat
- ✅ **Desain Modern** - UI responsif dan menarik

## 🛠️ Tech Stack

### Backend
- **Go 1.21+**
- **Gin Framework** - REST API
- **GORM** - ORM untuk database
- **JWT** - Autentikasi
- **Bcrypt** - Password hashing
- **MySQL/PostgreSQL** - Database

### Frontend
- **HTML5** - Struktur
- **CSS3** - Styling modern
- **JavaScript (Vanilla)** - Interaktivitas
- **Fetch API** - Komunikasi dengan backend

## 📦 Instalasi & Setup

### Prerequisites
- Go 1.21 atau lebih tinggi
- MySQL 5.7 atau PostgreSQL 10+
- Node.js (optional, untuk development tools)

### Backend Setup

1. **Clone/Download Project**
```bash
cd backend
```

2. **Install Dependencies**
```bash
go mod download
```

3. **Setup Database**
```sql
CREATE DATABASE skincare_db;
```

4. **Konfigurasi Environment**
Buat file `.env` berdasarkan `.env.example`:
```bash
cp .env.example .env
```

Edit file `.env`:
```
PORT=8080
GIN_MODE=debug

DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=skincare_db

JWT_SECRET=your-secret-key-here-change-this-in-production
```

5. **Jalankan Server**
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

### Frontend Setup

Frontend sudah tersedia di folder `frontend/` dan akan di-serve langsung oleh backend.

## 🚀 API Endpoints

### Authentication

**POST** `/api/auth/register`
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**POST** `/api/auth/login`
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

### User Profile

**GET** `/api/profile` (Requires Auth)
```
Headers: Authorization: Bearer <token>
```

### Skin Analysis

**POST** `/api/skin-analysis` (Requires Auth)
```json
{
  "skin_type": "berminyak",
  "skin_issues": ["jerawat", "kusam"],
  "sensitivity": "sedang"
}
```

**GET** `/api/history` (Requires Auth)
Mendapatkan semua riwayat analisis pengguna

**GET** `/api/history/:id` (Requires Auth)
Mendapatkan detail analisis spesifik

## 📝 Database Schema

### Users Table
```sql
CREATE TABLE users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at BIGINT,
  updated_at BIGINT
);
```

### Skin Analysis Table
```sql
CREATE TABLE skin_analysis (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  skin_type VARCHAR(50),
  skin_issues JSON,
  sensitivity VARCHAR(50),
  morning_routine JSON,
  night_routine JSON,
  created_at BIGINT,
  updated_at BIGINT,
  FOREIGN KEY(user_id) REFERENCES users(id)
);
```

## 🎨 Fitur UI

### Halaman Authentication
- **Login Page** - Form login dengan validasi
- **Register Page** - Form registrasi dengan konfirmasi password

### Halaman Dashboard
- **Profile Section** - Menampilkan data pengguna
- **Latest Analysis** - Analisis terbaru pengguna
- **History List** - Daftar semua analisis

### Halaman Analisis
- **Skin Type Selection** - Pilih jenis kulit
- **Issues Checkboxes** - Pilih masalah kulit
- **Sensitivity Dropdown** - Pilih tingkat sensitivitas
- **Result Display** - Tampilkan rutinitas pagi & malam

## 💡 Logika Rekomendasi

Sistem rekomendasi berdasarkan rule:

### Morning Routine
- Pembersih wajah sesuai jenis kulit
- Toner & essence berdasarkan jenis kulit
- Treatment khusus untuk masalah kulit
- Sunscreen sebagai langkah terakhir

### Night Routine
- Double cleanse (makeup remover + pembersih)
- Treatment aktif (retinol, acid) untuk malam
- Night moisturizer & sleeping mask
- Tips khusus berdasarkan sensitivitas

## 🔐 Keamanan

- Password di-hash menggunakan bcrypt
- JWT token untuk autentikasi
- CORS enabled untuk frontend
- Input validation di backend
- Protected routes dengan middleware

## 🧪 Testing API dengan cURL

```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'

# Get Profile (replace TOKEN dengan token dari login)
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer TOKEN"

# Create Analysis
curl -X POST http://localhost:8080/api/skin-analysis \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer TOKEN" \
  -d '{
    "skin_type": "berminyak",
    "skin_issues": ["jerawat", "kusam"],
    "sensitivity": "sedang"
  }'

# Get History
curl -X GET http://localhost:8080/api/history \
  -H "Authorization: Bearer TOKEN"
```

## 📱 Responsivitas

Aplikasi fully responsive untuk:
- Desktop (1200px+)
- Tablet (768px - 1199px)
- Mobile (< 768px)

## 🐛 Troubleshooting

### Database Connection Error
```
Error: Error 1045 (28000): Access denied for user 'root'@'localhost'
```
Solusi: Periksa konfigurasi `.env` (DB_USER, DB_PASSWORD, DB_HOST)

### CORS Error di Frontend
Pastikan backend sudah running dan endpoint benar di `dashboard.js`

### Token Invalid
Token expires setelah 24 jam. User perlu login ulang.

## 📚 Struktur Folder

```
skincare-app/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── .env.example
│   ├── config/
│   │   └── database.go
│   ├── models/
│   │   ├── user.go
│   │   └── skin_analysis.go
│   ├── controllers/
│   │   ├── auth.go
│   │   └── skincare.go
│   ├── middleware/
│   │   └── auth.go
│   └── routes/
│       └── routes.go
└── frontend/
    ├── index.html
    ├── login.html
    ├── register.html
    ├── dashboard.html
    ├── analysis.html
    └── assets/
        ├── style.css
        ├── auth.js
        ├── dashboard.js
        └── api.js
```

## 🚢 Deployment

### Local Development
```bash
cd backend
go run main.go
```

### Docker (Optional)
```dockerfile
FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o skincare-app

EXPOSE 8080
CMD ["./skincare-app"]
```

## 📝 License

MIT License - Silakan gunakan untuk keperluan apapun.

## 🤝 Kontribusi

Kontribusi welcome! Silakan buat pull request atau issue untuk improvement.

## ✨ Fitur Tambahan (Future)

- [ ] Rekomendasi produk spesifik
- [ ] Progress tracking skincare
- [ ] Photo upload untuk analisis visual
- [ ] Mobile app (React Native/Flutter)
- [ ] Email notifications
- [ ] Admin dashboard
- [ ] Social sharing
