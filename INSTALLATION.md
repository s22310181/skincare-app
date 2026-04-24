# Installation & Running Guide

## Persyaratan Sistem
- Go 1.21 atau lebih tinggi
- MySQL 5.7+ atau PostgreSQL 10+
- Browser modern
- Optional: Postman atau cURL untuk testing API

## Langkah-Langkah Instalasi

### 1. Setup Database

#### Menggunakan MySQL:
```bash
# Login ke MySQL
mysql -u root -p

# Jalankan script SQL
source database.sql

# Atau manual:
CREATE DATABASE skincare_db;
```

#### Menggunakan PostgreSQL:
```sql
CREATE DATABASE skincare_db;
```

### 2. Setup Backend

```bash
# Navigate ke folder backend
cd backend

# Download dependencies
go mod download

# Setup environment file
cp .env.example .env

# Edit .env dengan database credentials Anda
# Buka .env dan sesuaikan:
# DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, JWT_SECRET
```

### 3. Jalankan Backend

```bash
go run main.go
```

Output yang diharapkan:
```
Server starting on port 8080
```

Backend sekarang running di `http://localhost:8080`

### 4. Akses Frontend

Buka browser dan navigate ke:
- `http://localhost:8080/login` - Login page
- `http://localhost:8080/register` - Register page
- `http://localhost:8080/dashboard` - Dashboard (setelah login)

## Testing API

### 1. Register User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

Response:
```json
{
  "message": "User registered successfully",
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

Simpan token dari response untuk request berikutnya.

### 3. Get Profile
```bash
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 4. Create Skin Analysis
```bash
curl -X POST http://localhost:8080/api/skin-analysis \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "skin_type": "berminyak",
    "skin_issues": ["jerawat", "kusam"],
    "sensitivity": "sedang"
  }'
```

### 5. Get History
```bash
curl -X GET http://localhost:8080/api/history \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 6. Get Analysis Detail
```bash
curl -X GET http://localhost:8080/api/history/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Troubleshooting

### Error: Can't connect to database
**Solusi:**
1. Pastikan MySQL/PostgreSQL sudah running
2. Check database credentials di `.env`
3. Verifikasi database sudah dibuat

### Error: JWT_SECRET not set
**Solusi:**
1. Update `.env` dengan JWT_SECRET
2. Restart server

### CORS Error di Browser
**Solusi:**
- Backend sudah menghandle CORS
- Pastikan request dari origin yang benar
- Clear browser cache

### Frontend not loading
**Solusi:**
1. Pastikan backend running di port 8080
2. Check URL di browser: `http://localhost:8080`
3. Check browser console untuk errors

## Environment Variables Penjelasan

```env
PORT=8080                    # Port tempat server running
GIN_MODE=debug              # Mode: debug atau release
DB_USER=root                # Username database
DB_PASSWORD=password        # Password database
DB_HOST=localhost           # Host database
DB_PORT=3306                # Port database (MySQL: 3306, PostgreSQL: 5432)
DB_NAME=skincare_db         # Nama database
JWT_SECRET=key              # Secret key untuk JWT (generate yang aman!)
```

## Production Deployment

### Pre-deployment Checklist

1. **Security:**
   - [ ] Update JWT_SECRET dengan string yang panjang dan kompleks
   - [ ] Ganti DB_PASSWORD dengan password yang kuat
   - [ ] Gunakan HTTPS
   - [ ] Set GIN_MODE=release

2. **Database:**
   - [ ] Backup database sebelum deploy
   - [ ] Update database connection string
   - [ ] Run migration scripts

3. **Backend:**
   ```bash
   GIN_MODE=release go build -o skincare-app
   ./skincare-app
   ```

## File Penting

- `main.go` - Entry point aplikasi
- `config/database.go` - Database configuration
- `models/` - Data models
- `controllers/` - Business logic
- `routes/` - API routes
- `middleware/` - Auth middleware
- `frontend/` - HTML/CSS/JS files

## Performa Tips

1. Enable caching di frontend
2. Optimize database queries
3. Use CDN untuk static assets
4. Implement rate limiting di API
5. Monitor server logs

## Support

Jika ada masalah:
1. Check console/terminal output
2. Review logs
3. Verify all environment variables
4. Test API endpoints dengan Postman
