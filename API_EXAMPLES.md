# Library of Bears API Examples

## 1. Health Check
```bash
curl -X GET http://localhost:8000/healthcheck
```

## 2. User Registration
```bash
curl -X POST http://localhost:8000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword",
    "role": 1
  }'
```

## 3. User Login
```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```
**Note: Save the token from the response for authenticated requests**

## 4. Create Author (Protected)
```bash
curl -X POST http://localhost:8000/authors \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Jane Author"
  }'
```

## 5. Get All Authors (Public)
```bash
curl -X GET http://localhost:8000/authors
```

## 6. Create Series (Protected)
```bash
curl -X POST http://localhost:8000/series \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Amazing Series",
    "author_id": 1
  }'
```

## 7. Create Arts without File (Protected)
```bash
curl -X POST http://localhost:8000/arts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Beautiful Art",
    "description": "A wonderful piece of art",
    "author_id": 1,
    "series_id": 1,
    "source": "Original",
    "website": "https://example.com"
  }'
```

## 8. Create Arts with File Upload (Protected)
```bash
curl -X POST http://localhost:8000/arts/upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=Art with File" \
  -F "description=Beautiful uploaded artwork" \
  -F "author_id=1" \
  -F "series_id=1" \
  -F "source=Upload" \
  -F "website=https://example.com" \
  -F "file=@/path/to/your/image.jpg"
```

## 9. Get All Arts (Public)
```bash
curl -X GET http://localhost:8000/arts
```

## 10. Get Arts by Author (Public)
```bash
curl -X GET http://localhost:8000/arts/author/1
```

## 11. Update Arts (Protected)
```bash
curl -X PUT http://localhost:8000/arts/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Updated Art Title",
    "description": "Updated description",
    "author_id": 1,
    "series_id": 1
  }'
```

## 12. Delete Arts (Protected)
```bash
curl -X DELETE http://localhost:8000/arts/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Environment Setup
Before running these commands, make sure you have:

1. **Database running** (MySQL on port 3307)
2. **MinIO running** (on port 9000 for file uploads)
3. **Application running** (`go run .` or `./BearLibrary`)

You can use the provided docker-compose.yml:
```bash
docker-compose up -d mysql minio
```

## Response Format
All responses follow this structure:
```json
{
  "message": "success message",
  "data": {},
  "error": "error message if any"
}
```

Authentication responses include a `token` field:
```json
{
  "message": "Login successful",
  "user": {...},
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```