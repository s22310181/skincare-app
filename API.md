# 📋 API Documentation

## Base URL
```
http://localhost:8080/api
```

## Authentication
Semua endpoint (kecuali register & login) memerlukan JWT token di header:
```
Authorization: Bearer <token>
```

---

## 🔐 Authentication Endpoints

### 1. Register
**POST** `/auth/register`

**Request:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (201):**
```json
{
  "message": "User registered successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

**Error Response (400):**
```json
{
  "error": "User already exists"
}
```

---

### 2. Login
**POST** `/auth/login`

**Request:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

**Error Response (401):**
```json
{
  "error": "Invalid email or password"
}
```

---

## 👤 User Endpoints

### 3. Get Profile
**GET** `/profile`

**Headers:**
```
Authorization: Bearer <token>
```

**Success Response (200):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Error Response (401):**
```json
{
  "error": "Missing authorization header"
}
```

---

## 🧴 Skin Analysis Endpoints

### 4. Create Skin Analysis
**POST** `/skin-analysis`

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request:**
```json
{
  "skin_type": "berminyak",
  "skin_issues": ["jerawat", "kusam"],
  "sensitivity": "sedang"
}
```

**Parameters:**

| Parameter | Type | Required | Values |
|-----------|------|----------|--------|
| skin_type | string | Yes | `kering`, `berminyak`, `kombinasi`, `sensitif` |
| skin_issues | array | Yes | `jerawat`, `kusam`, `flek hitam`, `kerutan`, `kering`, `sensitif` |
| sensitivity | string | Yes | `rendah`, `sedang`, `tinggi` |

**Success Response (201):**
```json
{
  "message": "Skin analysis created successfully",
  "data": {
    "id": 1,
    "skin_type": "berminyak",
    "skin_issues": ["jerawat", "kusam"],
    "sensitivity": "sedang",
    "morning_routine": {
      "steps": [
        "1. Pembersih wajah - Cuci wajah dengan pembersih yang sesuai dengan jenis kulit",
        "2. Toner - Gunakan toner oil-control",
        ...
      ],
      "tips": "Pagi adalah waktu terbaik untuk menggunakan vitamin C dan sunscreen untuk perlindungan maksimal"
    },
    "night_routine": {
      "steps": [...],
      "tips": "..."
    },
    "created_at": 1234567890,
    "updated_at": 1234567890
  }
}
```

**Error Response (400):**
```json
{
  "error": "Validation error message"
}
```

---

### 5. Get Analysis History
**GET** `/history`

**Headers:**
```
Authorization: Bearer <token>
```

**Success Response (200):**
```json
{
  "data": [
    {
      "id": 2,
      "skin_type": "kombinasi",
      "skin_issues": ["kusam"],
      "sensitivity": "sedang",
      "morning_routine": {...},
      "night_routine": {...},
      "created_at": 1234567890,
      "updated_at": 1234567890
    },
    {
      "id": 1,
      "skin_type": "berminyak",
      "skin_issues": ["jerawat", "kusam"],
      "sensitivity": "sedang",
      "morning_routine": {...},
      "night_routine": {...},
      "created_at": 1234567800,
      "updated_at": 1234567800
    }
  ]
}
```

---

### 6. Get Analysis Detail
**GET** `/history/:id`

**Headers:**
```
Authorization: Bearer <token>
```

**Example:**
```
GET /history/1
```

**Success Response (200):**
```json
{
  "data": {
    "id": 1,
    "skin_type": "berminyak",
    "skin_issues": ["jerawat", "kusam"],
    "sensitivity": "sedang",
    "morning_routine": {
      "steps": [...],
      "tips": "..."
    },
    "night_routine": {
      "steps": [...],
      "tips": "..."
    },
    "created_at": 1234567890,
    "updated_at": 1234567890
  }
}
```

**Error Responses:**

404 - Analysis not found:
```json
{
  "error": "Analysis not found"
}
```

403 - Unauthorized:
```json
{
  "error": "Unauthorized"
}
```

---

## 🔍 Skin Type Guide

### Kulit Kering
- Karakteristik: Kulit terasa ketat, menggesek, mudah terkelupas
- Rekomendasi: Moisturizer kaya, essence hydrating

### Kulit Berminyak
- Karakteristik: Kilau berlebih, pori-pori besar, mudah berjerawat
- Rekomendasi: Toner oil-control, gel moisturizer

### Kulit Kombinasi
- Karakteristik: T-zone berminyak, pipi kering
- Rekomendasi: Produk balancing, ringan

### Kulit Sensitif
- Karakteristik: Mudah iritasi, kemerahan, reaksi terhadap produk
- Rekomendasi: Hypoallergenic, bebas fragrance

---

## 📊 Skin Issues

| Issue | Deskripsi |
|-------|-----------|
| Jerawat | Komedo, papul, pustul, nodul |
| Kusam | Kulit terlihat suram, tidak bersinar |
| Flek Hitam | Hiperpigmentasi, dark spots |
| Kerutan | Fine lines, wrinkles |
| Kering | Kering ekstrem, dehidrasi |
| Sensitif | Reaksi terhadap produk tertentu |

---

## 🔑 Token Format

JWT Token terdiri dari 3 parts:
1. **Header** - Algoritma & type
2. **Payload** - User data & claims
3. **Signature** - Tanda tangan

Example:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImpvaG5AZXhhbXBsZS5jb20iLCJleHAiOjE2OTQ2NzM1NTB9.
TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
```

Token expires setelah **24 jam**

---

## 🚨 Error Codes

| Code | Meaning | Solution |
|------|---------|----------|
| 200 | OK | Request berhasil |
| 201 | Created | Resource berhasil dibuat |
| 400 | Bad Request | Validasi input gagal |
| 401 | Unauthorized | Token invalid/expired atau login diperlukan |
| 403 | Forbidden | User tidak punya akses |
| 404 | Not Found | Resource tidak ditemukan |
| 500 | Server Error | Terjadi error di server |

---

## 📝 Request Examples dengan cURL

### Register
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com",
    "password": "securepassword123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane@example.com",
    "password": "securepassword123"
  }'
```

### Analyze Skin
```bash
curl -X POST http://localhost:8080/api/skin-analysis \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "skin_type": "kering",
    "skin_issues": ["kusam", "kering"],
    "sensitivity": "tinggi"
  }'
```

### Get History
```bash
curl -X GET http://localhost:8080/api/history \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 💡 Best Practices

1. **Token Storage**
   - Simpan token di `localStorage` untuk persistent login
   - Hapus token saat logout

2. **Error Handling**
   - Always check response status
   - Display user-friendly error messages
   - Log errors untuk debugging

3. **Security**
   - Gunakan HTTPS di production
   - Jangan expose JWT secret
   - Validate input di frontend dan backend
   - Use secure password practices

4. **Performance**
   - Cache responses ketika memungkinkan
   - Implement pagination untuk list besar
   - Optimize database queries

---

## 🔄 Workflow

1. User Register → Get Token
2. User Login → Get Token
3. User Access Protected Routes → Token di header
4. User Create Analysis → Save di database
5. User View History → Fetch dari database
6. User Logout → Remove Token

---

Untuk lebih detail, lihat `README.md` dan `INSTALLATION.md`
