# Japanese Courses API

A RESTful API service for a Japanese language learning platform built with Go and Echo framework.

## Features

- User Authentication (Register/Login) with JWT
- Multi-version API support (v1 and v2)
- Vocabulary Management
- Progress Tracking
- Topic and Lesson Organization
- Clean Architecture Implementation

## Tech Stack

- **Language:** Go
- **Framework:** Echo v4.11.4
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Migration Tool:** golang-migrate

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- golang-migrate (for database migrations)

## Project Structure

```
.
├── cmd/
├── internal/
│   ├── delivery/
│   │   └── api/
│   │       ├── v1/
│   │       └── v2/
│   ├── domain/
│   ├── infrastructure/
│   │   ├── auth/
│   │   └── database/
│   ├── models/
│   ├── repository/
│   └── usecase/
├── migrations/
└── main.go
```

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/JapaneseCourses.git
cd JapaneseCourses
```

2. Set up the database:
```bash
# Create PostgreSQL database
createdb japanese_courses

# Run migrations
migrate -database "postgres://postgres:your_password@localhost:5432/japanese_courses?sslmode=disable" -path migrations up
```

3. Configure the application:
   - Update database credentials in `main.go`
   - Update JWT secret key in router configuration

4. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

- **Register User (v2)**
  ```
  POST /api/v2/users/auth/register
  {
    "username": "string",
    "password": "string",
    "email": "string",
    "full_name": "string"
  }
  ```

- **Login (v2)**
  ```
  POST /api/v2/users/auth/login
  {
    "username": "string",
    "password": "string"
  }
  ```

### Vocabulary

- **Get All Vocabulary**
  ```
  GET /api/v2/vocabulary
  ```

- **Get Vocabulary by ID**
  ```
  GET /api/v2/vocabulary/:id
  ```

- **Get Learned Vocabulary**
  ```
  GET /api/v2/vocabulary-progress/learned
  Authorization: Bearer <token>
  ```

## Database Schema

The application uses the following main tables:
- users
- vocabulary
- topics
- lessons
- vocabulary_progress
- tests
- test_results

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. # Japanese-Courses
# Japanese-Courses
