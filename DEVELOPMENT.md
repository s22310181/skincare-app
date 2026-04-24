# 📝 Development Notes & Code Guide

## Architecture Overview

### Monolithic Architecture
```
┌─────────────────────────────────────┐
│      Frontend (HTML/CSS/JS)         │
├─────────────────────────────────────┤
│      Gin HTTP Server (Go)            │
│  - Routes                            │
│  - Controllers                       │
│  - Models                            │
│  - Middleware                        │
├─────────────────────────────────────┤
│      Database (MySQL/PostgreSQL)     │
└─────────────────────────────────────┘
```

### Component Relationships
```
Frontend Request
    ↓
Gin Router (match route)
    ↓
Middleware (auth validation)
    ↓
Controller (business logic)
    ↓
Model (database operations)
    ↓
Database
    ↓
Response JSON
    ↓
Frontend Display
```

---

## Code Organization

### Main Entry Point: main.go

```go
// Responsibility: Application startup & configuration
- Load environment variables
- Initialize database connection
- Setup Gin router
- Configure CORS
- Setup routes
- Start HTTP server
```

### Config: config/database.go

```go
// Responsibility: Database initialization
- Build DSN connection string
- Open database connection
- Run auto-migrations
- Handle connection errors
```

### Models: models/

```go
// users.go
User struct {
  ID, Name, Email, Password, timestamps
}

// Functions: GetUserByEmail, GetUserByID, CreateUser

// skin_analysis.go
SkinAnalysis struct {
  ID, UserID, SkinType, Issues, Sensitivity, Routines
}

// Functions: GetAnalysisByID, GetAnalysisByUserID, CreateAnalysis
```

### Controllers: controllers/

```go
// auth.go
- Register(c *gin.Context)    // Handle user signup
- Login(c *gin.Context)       // Handle user login
- GetProfile(c *gin.Context)  // Return user data

// skincare.go
- AnalyzeSkin(c *gin.Context)        // Create analysis
- GetHistory(c *gin.Context)         // List analyses
- GetAnalysisDetail(c *gin.Context)  // Get single analysis
- Helper functions for recommendations
```

### Middleware: middleware/

```go
// auth.go
- AuthMiddleware() gin.HandlerFunc    // Verify JWT token
- GetUserID(c *gin.Context) uint      // Extract user ID from token
- Claims struct                        // JWT payload structure
```

### Routes: routes/

```go
// routes.go
- SetupRoutes(router, db)             // Configure all API routes
  - Public: /api/auth/*
  - Protected: /api/profile, /api/skin-analysis, /api/history
```

---

## Frontend Architecture

### Page Structure
```
index.html (landing)
  ├─ Redirects to login if not authenticated
  └─ Shows home info

login.html
  ├─ Email & password form
  ├─ Links to register
  └─ Stores token on success

register.html
  ├─ Name, email, password form
  ├─ Password confirmation
  ├─ Links to login
  └─ Stores token on success

dashboard.html (protected)
  ├─ Navigation bar
  ├─ Dashboard page
  │  ├─ Profile display
  │  ├─ Latest analysis
  │  └─ History list
  └─ Analysis page
     ├─ Skin analysis form
     ├─ Result display
     └─ History

assets/
  ├─ style.css       - Global styling
  ├─ auth.js         - Auth page handlers
  ├─ dashboard.js    - Dashboard logic
  └─ api.js          - API utilities
```

### Frontend Data Flow
```
User Input (Form)
    ↓
JavaScript Validation
    ↓
API Call (fetch)
    ↓
Backend Processing
    ↓
JSON Response
    ↓
Update DOM
    ↓
User Feedback (message, redirect)
```

---

## Database Design Decisions

### Users Table
- **Why indexed email**: Quick lookup during login
- **Why hashed password**: Security - passwords never in plain text
- **Why timestamps**: Track user lifecycle

### Skin Analysis Table
- **Why JSON for issues**: Flexible storage for multiple issues
- **Why JSON for routines**: Flexible structure for steps & tips
- **Why user_id foreign key**: Ensure data integrity & enable cascade delete
- **Why indexed user_id**: Fast history queries
- **Why indexed created_at**: Sort history efficiently

### No separate tables for:
- Issues (stored as JSON array)
- Routines (stored as JSON object)
- Recommendations (generated on-the-fly)

This keeps schema simple while maintaining flexibility.

---

## API Design Principles

### RESTful Endpoints
```
POST   /auth/register        - Create new user
POST   /auth/login          - User authentication
GET    /profile             - Get user profile
POST   /skin-analysis       - Create analysis
GET    /history             - List analyses
GET    /history/:id         - Get analysis detail
```

### Status Codes
- 200: OK (GET, successful operations)
- 201: Created (POST successful)
- 400: Bad Request (validation error)
- 401: Unauthorized (auth required or invalid)
- 403: Forbidden (user not allowed)
- 404: Not Found (resource missing)
- 500: Server Error (unexpected error)

### Error Response Format
```json
{
  "error": "Human readable message"
}
```

### Success Response Format
```json
{
  "message": "Optional status message",
  "data": { /* response data */ },
  "token": "JWT token (auth endpoints only)"
}
```

---

## Security Implementation

### Password Security
```go
// Generate hash
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Compare password
err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
```

### JWT Implementation
```go
// Create token
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString([]byte(JWT_SECRET))

// Validate token
token, _ := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
  return []byte(JWT_SECRET), nil
})
```

### CORS Configuration
```go
// Allow all origins (development only!)
c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

// For production, restrict to specific domains
```

---

## Recommendation Engine Logic

### Decision Tree
```
Input: skin_type, skin_issues, sensitivity

1. Base routine (skin_type specific)
   ├─ Pembersih wajah
   ├─ Toner
   └─ Moisturizer

2. Add specific treatments
   ├─ If has "jerawat" → add acne treatment
   ├─ If has "kusam" → add vitamin C / retinol
   ├─ If has "flek hitam" → add spot treatment
   └─ If has "kerutan" → add anti-aging serum

3. Adjust by sensitivity
   ├─ tinggi → gentle, hypoallergenic, fragrance-free
   ├─ sedang → balanced approach
   └─ rendah → can use more active ingredients

4. Add final steps
   ├─ Morning: Sunscreen SPF 30+
   └─ Night: Optional sleeping mask/heavy cream

Output: morning_routine & night_routine with steps & tips
```

### Implementation
Functions in `controllers/skincare.go`:
- `generateMorningRoutine()` - Creates morning steps
- `generateNightRoutine()` - Creates night steps
- `contains()` - Helper to check if issue exists

---

## Common Development Tasks

### Adding New Endpoint

1. **Add Route** (routes/routes.go)
```go
protectedRoutes.GET("/new-endpoint", controllers.NewHandler)
```

2. **Create Controller** (controllers/skincare.go)
```go
func NewHandler(c *gin.Context) {
  // Get user ID from middleware
  userID := middleware.GetUserID(c)
  
  // Process request
  db := c.MustGet("db").(*gorm.DB)
  
  // Return response
  c.JSON(http.StatusOK, gin.H{"data": data})
}
```

3. **Test with cURL**
```bash
curl -X GET http://localhost:8080/api/new-endpoint \
  -H "Authorization: Bearer TOKEN"
```

### Adding New Database Model

1. **Create Model** (models/new_model.go)
```go
type NewModel struct {
  ID    uint
  Name  string
}

func CreateNew(db *gorm.DB, model *NewModel) error {
  return db.Create(model).Error
}
```

2. **Auto-migrate** (config/database.go)
```go
err = db.AutoMigrate(&models.NewModel{})
```

3. **Use in Controller**
```go
models.CreateNew(db, &newModel)
```

### Updating UI

1. **Add HTML** (frontend/*.html)
```html
<form id="newForm">
  <input type="text" id="newInput">
  <button type="submit">Submit</button>
</form>
```

2. **Add JavaScript** (frontend/assets/*.js)
```javascript
document.getElementById('newForm').addEventListener('submit', async (e) => {
  e.preventDefault();
  // Handle form submission
});
```

3. **Add Styling** (frontend/assets/style.css)
```css
#newForm {
  /* Your styles */
}
```

---

## Performance Optimization Tips

### Database Queries
```go
// ❌ Bad: N+1 queries
for _, user := range users {
  var analyses []SkinAnalysis
  db.Where("user_id = ?", user.ID).Find(&analyses)
}

// ✅ Good: Eager loading
db.Preload("Analysis").Find(&users)
```

### API Responses
```go
// ❌ Bad: Unnecessary data
c.JSON(200, user) // Returns password hash!

// ✅ Good: Only needed fields
c.JSON(200, UserResponse{
  ID:    user.ID,
  Name:  user.Name,
  Email: user.Email,
})
```

### Frontend Caching
```javascript
// ❌ Bad: Always fetch
fetch('/api/history')

// ✅ Good: Cache where possible
const cached = localStorage.getItem('userHistory')
if (cached && !expired) {
  return JSON.parse(cached)
}
fetch('/api/history')
```

---

## Testing Strategies

### Manual API Testing
```bash
# 1. Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@test.com","password":"123456"}'

# 2. Copy token from response

# 3. Use token in protected routes
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer TOKEN"
```

### Frontend Testing
```javascript
// Test in browser console
const token = localStorage.getItem('token')
fetch('http://localhost:8080/api/profile', {
  headers: { 'Authorization': `Bearer ${token}` }
}).then(r => r.json()).then(console.log)
```

### Database Testing
```sql
-- Check users table
SELECT * FROM users WHERE email = 'test@test.com';

-- Check analyses
SELECT * FROM skin_analysis WHERE user_id = 1;

-- Verify foreign key
SELECT * FROM skin_analysis s 
JOIN users u ON s.user_id = u.id 
WHERE u.id = 1;
```

---

## Debugging Tips

### Go Debugging
```bash
# Add debug print
fmt.Printf("Debug: %+v\n", variable)

# Check compilation
go build

# Run with race detector
go run -race main.go
```

### Browser Debugging
```javascript
// Console logs
console.log('Debug:', variable)
console.error('Error:', error)
console.table(data)

// Breakpoints in DevTools
debugger; // Pauses execution

// Network inspection
// Open DevTools → Network tab → watch requests
```

### Database Debugging
```sql
-- Check connection
SELECT 1;

-- View table structure
DESC users;

-- Check indexes
SHOW INDEX FROM users;

-- Monitor slow queries
-- Set: SET GLOBAL slow_query_log = 'ON';
```

---

## Code Style Guidelines

### Go Convention
```go
// Package names: lowercase, no underscores
package models

// Function names: PascalCase for exported
func GetUserByID(db *gorm.DB, id uint) *User {}

// Private functions: camelCase
func privateHelper() {}

// Constants: UPPER_CASE
const MaxPasswordLength = 255

// Comments on exported items
// GetUserByID retrieves a user from database
func GetUserByID(...)
```

### JavaScript Convention
```javascript
// Function names: camelCase
function getUserData() {}

// Class names: PascalCase
class UserManager {}

// Constants: UPPER_CASE
const API_BASE_URL = 'http://localhost:8080/api'

// Variables: camelCase
let userData = {}
```

### CSS Convention
```css
/* Class names: kebab-case */
.user-profile {
  /* Properties: alphabetically organized */
  color: #333;
  padding: 10px;
}

/* IDs: camelCase */
#userForm {}
```

---

## Deployment Considerations

### Environment-Specific Config
```
Development:
  GIN_MODE=debug
  LOG_LEVEL=debug
  ALLOW_CORS=*
  
Production:
  GIN_MODE=release
  LOG_LEVEL=warn
  ALLOW_CORS=https://yourdomain.com
  DB uses RDS or managed service
```

### Secrets Management
```
❌ Don't:
- Commit .env with real secrets
- Store secrets in code
- Use weak JWT_SECRET

✅ Do:
- Use .env.example as template
- Use environment variable injection
- Use strong random JWT_SECRET
- Use managed secrets (AWS Secrets Manager, etc.)
```

### Monitoring Essentials
```
Log:
  - API errors
  - Database errors
  - Authentication failures
  - Slow queries (> 100ms)

Monitor:
  - Server uptime
  - Response times
  - Database performance
  - Error rates
  - User growth
```

---

## Future Enhancements

### Short Term
- [ ] Add pagination to history
- [ ] Add product recommendations
- [ ] Add user profile editing
- [ ] Add analysis deletion

### Medium Term
- [ ] Add progress tracking (photos)
- [ ] Add product marketplace integration
- [ ] Add email notifications
- [ ] Add admin dashboard

### Long Term
- [ ] Mobile app (React Native)
- [ ] AI/ML recommendations
- [ ] Community features
- [ ] Payment processing

---

## Resources & References

### Go
- https://golang.org/doc/
- https://gin-gonic.com/docs/
- https://gorm.io/docs/

### Database
- https://dev.mysql.com/doc/
- https://www.postgresql.org/docs/

### Frontend
- https://developer.mozilla.org/
- https://www.w3schools.com/

### Security
- https://owasp.org/
- https://cheatsheetseries.owasp.org/

---

Happy coding! 🚀

Last Updated: April 2026
Version: 1.0.0
