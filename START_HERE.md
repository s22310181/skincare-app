# 🚀 GETTING STARTED - Read This First!

## 👋 Welcome!

Anda baru saja menerima **complete fullstack web application** bernama **SkinCare Routine App**.

Aplikasi ini adalah production-ready dan bisa langsung dijalankan.

## ⏱️ 5 Minute Quick Start

### Step 1: Setup Database (2 menit)

```bash
# Option A: Jika punya MySQL installed
mysql -u root -p < database.sql

# Option B: Jika nggak punya MySQL
docker run -d \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=skincare_db \
  mysql:8.0

# Wait 30 seconds untuk MySQL startup
```

### Step 2: Setup & Run Backend (3 menit)

```bash
cd backend

# Copy environment template
cp .env.example .env

# Edit .env - update DB_PASSWORD sesuai MySQL Anda
# Buka: .env
# Ubah: DB_PASSWORD=password (sesuaikan dengan password MySQL)
# Save file

# Download dependencies
go mod download

# Start server
go run main.go
```

Jika berhasil, Anda akan lihat:
```
Server starting on port 8080
```

### Step 3: Open Browser (Instant)

```
http://localhost:8080/login
```

## ✨ Try It Out!

1. **Click** "Belum punya akun? Daftar sekarang"
2. **Fill** form: Name, Email, Password
3. **Click** "Daftar"
4. **You're logged in!** Redirect ke dashboard
5. **Click** "Analisis"
6. **Select:**
   - Jenis Kulit: Berminyak
   - Masalah Kulit: Jerawat ✓
   - Sensitivitas: Sedang
7. **Click** "Analisis Kulit Saya"
8. **See** Morning & Night Routine recommendations!

## 📚 Documentation Files

| File | Waktu | Isi |
|------|-------|-----|
| **QUICKSTART.md** | 10 min | Setup guide lengkap |
| **README.md** | 15 min | Feature overview |
| **API.md** | 20 min | API endpoints |
| **INSTALLATION.md** | 30 min | Detailed setup |
| **FEATURES.md** | 20 min | Feature showcase |
| **DEVELOPMENT.md** | 1 hour | Code explanation |
| **PROJECT_SUMMARY.md** | 30 min | Architecture |

**👉 Next:** Baca **QUICKSTART.md** untuk setup detail

## 🔧 Common Issues & Fixes

### ❌ "Connection refused" - MySQL Error

**Problem:** Database connection failed

**Solution:**
```bash
# Check if MySQL running
mysql -u root -p

# If not, start MySQL or use Docker:
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=skincare_db mysql:8.0
```

### ❌ "Port 8080 in use" Error

**Problem:** Port sudah dipakai

**Solution:**
```bash
# Use different port
PORT=9090 go run main.go

# Then access: http://localhost:9090
```

### ❌ "GORM" Error atau dependency issue

**Problem:** Go modules error

**Solution:**
```bash
cd backend
go clean -modcache
go mod download
go run main.go
```

### ❌ "No such file" - Frontend pages missing

**Problem:** Frontend tidak ter-serve

**Solution:**
- Pastikan backend running dari root project
- URL harus: http://localhost:8080 (bukan localhost:8080/backend)

## 🎯 Project Structure Overview

```
Final-Backend/
├── 📖 Documentation (README, QUICKSTART, API, etc)
├── 🔴 backend/ (Go REST API)
│   ├── main.go (entry point)
│   ├── .env (configuration)
│   ├── config/ (database)
│   ├── models/ (data models)
│   ├── controllers/ (business logic)
│   ├── middleware/ (authentication)
│   └── routes/ (API routes)
├── 🔵 frontend/ (HTML/CSS/JS)
│   ├── index.html, login.html, register.html
│   ├── dashboard.html, analysis.html
│   └── assets/ (CSS, JavaScript)
└── 🗄️ database.sql (SQL schema)
```

## 🔐 Security Settings (Important!)

⚠️ **Defaults untuk development saja!**

### .env defaults:
```
JWT_SECRET=your-secret-key-here-change-this-in-production-12345
DB_PASSWORD=password
GIN_MODE=debug
```

### ✅ Untuk production:
1. Change JWT_SECRET ke random string yang panjang
2. Change DB_PASSWORD ke password yang kuat
3. Set GIN_MODE=release
4. Enable HTTPS
5. Use environment variable injection

Lihat **INSTALLATION.md** untuk detail.

## 🧪 Test It Immediately

### Via Browser (Recommended)
1. Go to http://localhost:8080/register
2. Fill form
3. Click register
4. Automatically logged in to dashboard
5. Click "Analisis"
6. Fill form & submit
7. See recommendations!

### Via API (cURL)

```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'

# Copy token from response
# Then test protected endpoint:

curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## 📞 Need Help?

1. **Setup issue?** → Read QUICKSTART.md
2. **API question?** → Read API.md  
3. **Code explanation?** → Read DEVELOPMENT.md
4. **Feature question?** → Read FEATURES.md
5. **Architecture?** → Read PROJECT_SUMMARY.md

## 🎓 Learning Path

### Beginner
1. Follow 5-minute quickstart above
2. Test app in browser
3. Read README.md

### Intermediate  
1. Read API.md
2. Test API with cURL
3. Read FEATURES.md
4. Explore code in backend/

### Advanced
1. Read DEVELOPMENT.md
2. Read code comments
3. Modify recommendation logic
4. Add new features

## 🚀 What's Included

✅ **Complete Backend** - Ready to run
✅ **Complete Frontend** - Ready to use
✅ **Database Schema** - Ready to import
✅ **API Documentation** - Ready to test
✅ **Deployment Files** - Docker ready
✅ **Code Documentation** - Well commented

## ⚡ Tech Stack

- **Backend:** Go + Gin Framework + GORM
- **Frontend:** HTML5 + CSS3 + Vanilla JavaScript
- **Database:** MySQL (PostgreSQL compatible)
- **Authentication:** JWT + Bcrypt
- **Deployment:** Docker + Docker Compose

## 🎊 You're Ready!

```bash
# 3 commands to get running:
mysql -u root -p < database.sql
cd backend && cp .env.example .env
go run main.go

# Then open: http://localhost:8080/login
```

That's it! 🎉

## 📋 Checklist

- [ ] MySQL ready (or Docker running)
- [ ] Cloned/downloaded project
- [ ] Navigated to backend folder
- [ ] Created .env file
- [ ] Updated DB_PASSWORD in .env
- [ ] Run: go run main.go
- [ ] Opened: http://localhost:8080/login
- [ ] Registered new user
- [ ] Tested skin analysis
- [ ] See recommendations

✅ **All done!** Now explore the code and documentation.

## 🔗 Quick Links

```
📄 Quick Setup:      QUICKSTART.md
📖 Full Overview:    README.md
🔗 API Reference:    API.md
🛠️  Development:      DEVELOPMENT.md
📊 Features:         FEATURES.md
✅ Testing:          CHECKLIST.md
🗂️  File Structure:   FILE_STRUCTURE.md
```

---

**Next Steps:**

1. ✅ Get app running (5 min above)
2. 📖 Read QUICKSTART.md for more detail
3. 🧪 Test all features in browser
4. 🔗 Read API.md if building mobile app
5. 💻 Read DEVELOPMENT.md to understand code
6. 🚀 Customize & deploy!

---

**Happy coding! You got this! 🚀✨**
