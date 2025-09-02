# Library of Bears

### Clean Architecture REST API for organizing arts, authors, and series

This project implements a clean architecture REST API with complete CRUD operations, JWT authentication, and file upload support using MinIO.

## Features

- **Clean Architecture**: Domain, UseCase, Repository, and Delivery layers
- **Complete CRUD APIs**: For Users, Authors, Series, and Arts
- **JWT Authentication**: Secure token-based authentication
- **File Upload**: Support for file uploads using MinIO storage
- **Password Security**: bcrypt password hashing
- **Proper Error Handling**: Comprehensive error responses
- **CORS Support**: Cross-origin resource sharing enabled

## Architecture

```
├── Models/                 # Domain entities
├── internal/
│   ├── domain/
│   │   └── repository/     # Repository interfaces
│   ├── infrastructure/
│   │   ├── repository/     # Repository implementations
│   │   └── service/        # External services (MinIO)
│   ├── usecase/           # Business logic
│   ├── delivery/
│   │   └── http/          # HTTP handlers
│   └── container/         # Dependency injection
├── Config/                # Configuration
├── Helper/                # Utilities (JWT)
└── Route/                 # API routes
```

## Environment Variables

```env
# Database
DB_HOST=localhost
DB_USER=root
DB_PORT=3307
DB_NAME=library_of_bears

# JWT
SECRET=your-secret-key

# MinIO
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_USE_SSL=false
```

## API Endpoints

### Authentication
- `POST /auth/register` - User registration
- `POST /auth/login` - User login

### Users (Protected)
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

### Authors
- `GET /authors` - Get all authors (public)
- `GET /authors/:id` - Get author by ID (public)
- `POST /authors` - Create author (protected)
- `PUT /authors/:id` - Update author (protected)
- `DELETE /authors/:id` - Delete author (protected)

### Series
- `GET /series` - Get all series (public)
- `GET /series/:id` - Get series by ID (public)
- `GET /series/author/:authorId` - Get series by author (public)
- `POST /series` - Create series (protected)
- `PUT /series/:id` - Update series (protected)
- `DELETE /series/:id` - Delete series (protected)

### Arts
- `GET /arts` - Get all arts (public)
- `GET /arts/:id` - Get arts by ID (public)
- `GET /arts/author/:authorId` - Get arts by author (public)
- `GET /arts/series/:seriesId` - Get arts by series (public)
- `POST /arts` - Create arts (protected)
- `POST /arts/upload` - Create arts with file upload (protected)
- `PUT /arts/:id` - Update arts (protected)
- `PUT /arts/:id/upload` - Update arts with file upload (protected)
- `DELETE /arts/:id` - Delete arts (protected)

### Utility
- `GET /healthcheck` - Health check endpoint

## Usage Examples

### Register a User
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

### Login
```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

### Create Author (Protected)
```bash
curl -X POST http://localhost:8000/authors \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "John Author"
  }'
```

### Upload Arts with File (Protected)
```bash
curl -X POST http://localhost:8000/arts/upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=My Art Piece" \
  -F "description=Beautiful artwork" \
  -F "author_id=1" \
  -F "file=@/path/to/image.jpg"
```

## Development

### Prerequisites
- Go 1.21+
- MySQL database
- MinIO server (for file uploads)

### Setup
1. Clone the repository
2. Set environment variables
3. Run `go mod tidy`
4. Start MySQL and MinIO servers
5. Run `go run .`

### Testing
```bash
go test -v
```

### Building
```bash
go build .
```

## MinIO Setup

To set up MinIO locally:

```bash
# Download and run MinIO
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server ~/minio-data --console-address ":9001"
```

Access MinIO console at http://localhost:9001 with credentials:
- Username: minioadmin
- Password: minioadmin