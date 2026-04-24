# 🚀 Quick Start Guide

Panduan cepat untuk memulai SkinCare Routine App.

## 1️⃣ Persiapan Awal (5 menit)

### Install Required Tools
- **Go**: https://golang.org/dl/ (versi 1.21+)
- **MySQL**: https://dev.mysql.com/downloads/mysql/ atau gunakan Docker
- **Git**: https://git-scm.com/

### Clone Project (Jika belum)
```bash
cd Final-Backend
```

## 2️⃣ Setup Database (5 menit)

### Option A: MySQL Native
```bash
# Login ke MySQL
mysql -u root -p

# Jalankan SQL
source database.sql

# Verifikasi
USE skincare_db;
SHOW TABLES;
```

### Option B: Docker
```bash
docker-compose up -d mysql

# Wait untuk mysql ready, kemudian:
docker exec -i skincare_mysql mysql -u root -ppassword skincare_db < database.sql
```

## 3️⃣ Setup Backend (3 menit)

```bash
cd backend

# Setup .env
cp .env.example .env

# Edit .env (pastikan DB credentials sesuai)
# Minimal: DB_PASSWORD harus sesuai dengan MySQL Anda

# Download dependencies
go mod download

# Jalankan server
go run main.go
```

Output:
```
Server starting on port 8080
```

✅ Backend running!

## 4️⃣ Akses Aplikasi

Buka browser dan navigasi:

| Halaman | URL |
|---------|-----|
| **Home** | http://localhost:8080 |
| **Login** | http://localhost:8080/login |
| **Register** | http://localhost:8080/register |
| **Dashboard** | http://localhost:8080/dashboard |

## 5️⃣ Test Aplikasi (2 menit)

### User Flow:
1. **Register** → http://localhost:8080/register
   - Name: `Test User`
   - Email: `test@example.com`
   - Password: `password123`

2. **Login** → http://localhost:8080/login
   - Email: `test@example.com`
   - Password: `password123`

3. **Dashboard** → View profile & riwayat

4. **Analisis** → Buat skincare analysis:
   - Skin Type: `Berminyak`
   - Issues: `Jerawat`, `Kusam`
   - Sensitivity: `Sedang`

5. **Lihat Hasil** → Rutinitas pagi & malam

## 🐳 Quick Start dengan Docker (2 menit)

```bash
# Build dan run semua service
docker-compose up --build

# Service akan running:
# - MySQL: localhost:3306
# - Backend: http://localhost:8080
```

Stop services:
```bash
docker-compose down
```

## 🔧 Konfigurasi Common Issues

### Database Connection Error
```
Error: dial tcp localhost:3306: connection refused
```

**Solusi:**
```bash
# Check jika MySQL running
mysql -u root -p

# Atau gunakan Docker
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=skincare_db mysql:8.0
```

### Port 8080 Already in Use
```bash
# Find process using port 8080
netstat -ano | findstr :8080

# Kill process (Windows)
taskkill /PID <PID> /F

# Atau gunakan port berbeda
PORT=9090 go run main.go
```

### Go Modules Issues
```bash
# Clear cache
go clean -modcache

# Re-download
go mod download
```

## 📱 Frontend Features

### Registrasi & Login
- ✅ Validasi email
- ✅ Password hashing
- ✅ JWT authentication
- ✅ Session persistence

### Dashboard
- ✅ Profile display
- ✅ Analysis history
- ✅ Latest analysis
- ✅ Logout button

### Skin Analysis
- ✅ Jenis kulit (4 opsi)
- ✅ Masalah kulit (6 opsi)
- ✅ Tingkat sensitivitas (3 level)
- ✅ Smart recommendations

### Routine Recommendations
- ✅ Morning routine dengan 6-7 langkah
- ✅ Night routine dengan 6-7 langkah
- ✅ Tips khusus per sensitivitas
- ✅ Saran berbasis masalah kulit

## 📚 Struktur Folder Quick Reference

```
Final-Backend/
├── backend/              # 🔴 Go backend
│   ├── main.go          # Entry point
│   ├── go.mod           # Dependencies
│   ├── .env             # Configuration (jangan share!)
│   ├── config/          # Database config
│   ├── models/          # Data models
│   ├── controllers/      # Business logic
│   ├── middleware/       # Auth middleware
│   └── routes/          # API routes
├── frontend/            # 🔵 Frontend files
│   ├── login.html       # Login page
│   ├── register.html    # Register page
│   ├── dashboard.html   # Dashboard
│   └── assets/          # CSS, JS
├── database.sql         # Database schema
├── docker-compose.yml   # Docker config
└── README.md           # Full documentation
```

## 🔐 Security Reminders

⚠️ **Development Defaults:**
- JWT_SECRET: `your-secret-key-here` (JANGAN di-production!)
- DB_PASSWORD: `password` (JANGAN di-production!)
- GIN_MODE: `debug` (JANGAN di-production!)

✅ **Production Checklist:**
- [ ] Generate strong JWT_SECRET
- [ ] Set strong DB_PASSWORD
- [ ] Set `GIN_MODE=release`
- [ ] Use HTTPS
- [ ] Remove debug logs
- [ ] Database backups enabled

## 🧪 Testing API dengan Postman

1. **Import Collection** (optional):
   - Create new request
   - Set method & URL
   - Add headers/body

2. **Register:**
   ```
   POST http://localhost:8080/api/auth/register
   Body: {
     "name": "Test User",
     "email": "test@example.com",
     "password": "password123"
   }
   ```

3. **Login & Copy Token:**
   ```
   POST http://localhost:8080/api/auth/login
   Body: {
     "email": "test@example.com",
     "password": "password123"
   }
   ```

4. **Create Analysis:**
   ```
   POST http://localhost:8080/api/skin-analysis
   Headers: Authorization: Bearer <TOKEN>
   Body: {
     "skin_type": "berminyak",
     "skin_issues": ["jerawat"],
     "sensitivity": "sedang"
   }
   ```

## 📖 Full Documentation

Untuk info lebih detail, baca:
- 📘 [README.md](./README.md) - Overview & features
- 🔗 [API.md](./API.md) - API endpoints & examples
- 📋 [INSTALLATION.md](./INSTALLATION.md) - Detailed setup

## 🎯 Next Steps

1. ✅ Run aplikasi
2. ✅ Register user baru
3. ✅ Create skin analysis
4. ✅ Lihat recommendations
5. 🔧 Customize logic di `controllers/skincare.go`
6. 🎨 Customize UI di `frontend/assets/style.css`
7. 🚀 Deploy!

## 💬 Tips & Tricks

- **Hardcoding test token?** Copy dari login response
- **Database corrupted?** Run `database.sql` ulang
- **Mau ubah recommendation logic?** Edit `generateMorningRoutine()` & `generateNightRoutine()`
- **Mau add field baru?** Update `models/`, `controllers/`, dan `frontend/`

## 🆘 Need Help?

1. Check console output untuk error messages
2. Verify .env configuration
3. Check database connection
4. Review API.md untuk endpoint format
5. Use browser DevTools untuk frontend debugging

---

**Ready to go!** 🚀

Jika ada pertanyaan atau issue, check documentation files atau debug di code. Happy coding! 💻✨
