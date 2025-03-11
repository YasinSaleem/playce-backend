# Playce Backend

A Go-based microservices backend built with the Echo framework, featuring user authentication and post management.

## Project Overview

This project consists of two microservices:

1. **User Service**: Handles user authentication, registration, and password management
2. **Post Service**: Manages posts creation and retrieval

The services are built using:
- Go programming language
- Echo web framework
- PostgreSQL database
- JWT for authentication

## Prerequisites

- Go 1.18 or later
- PostgreSQL
- Git

## Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd se_backend
```

### 2. Database Setup

Make sure PostgreSQL is running and create a database named `playce`:

```bash
psql -U postgres -c "CREATE DATABASE playce;"
```

The default database configuration in `config.json` is:

```json
{
    "database": {
        "url": "host=localhost user=postgres dbname=playce sslmode=disable password=postgres"
    }
}
```

Modify this if your PostgreSQL setup is different.

### 3. Build and Run the Services

#### User Service

```bash
cd user_service
go build
./user_service
```

Or run without building:

```bash
cd user_service
go run main.go
```

The user service will start on port 8080.

#### Post Service

```bash
cd post_service
go build
./post_service
```

Or run without building:

```bash
cd post_service
go run main.go
```

The post service will start on port 8081.

## API Documentation

### User Service (http://localhost:8080)

#### 1. Register a New User

**Endpoint**: `POST /user/signup`

**Request**:
```bash
curl -X POST http://localhost:8080/user/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Response**:
```json
{
  "message": "User registered successfully"
}
```

#### 2. Sign In

**Endpoint**: `POST /user/signin`

**Request**:
```bash
curl -X POST http://localhost:8080/user/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Response**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 3. Forgot Password

**Endpoint**: `POST /user/forgot-password`

**Request**:
```bash
curl -X POST http://localhost:8080/user/forgot-password \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com"
  }'
```

**Response**:
```json
{
  "message": "If your email is registered, you will receive a reset link",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

Note: In a production environment, the token would be sent via email rather than returned in the response.

#### 4. Reset Password

**Endpoint**: `POST /user/reset-password`

**Request**:
```bash
curl -X POST http://localhost:8080/user/reset-password \
  -H "Content-Type: application/json" \
  -d '{
    "token": "RESET_TOKEN_FROM_FORGOT_PASSWORD_RESPONSE",
    "password": "newpassword123"
  }'
```

**Response**:
```json
{
  "message": "Password reset successfully"
}
```

### Post Service (http://localhost:8081)

#### 1. Get All Posts

**Endpoint**: `GET /posts`

**Request**:
```bash
curl -X GET http://localhost:8081/posts
```

**Response**:
```json
[
  {
    "ID": 1,
    "CreatedAt": "2025-03-11T21:46:43.736069+05:30",
    "UpdatedAt": "2025-03-11T21:46:43.736069+05:30",
    "DeletedAt": null,
    "body": "",
    "user_id": 0
  }
]
```

#### 2. Get User Posts

**Endpoint**: `GET /posts/:userId`

**Request**:
```bash
curl -X GET http://localhost:8081/posts/1
```

**Response**:
```json
[
  {
    "ID": 1,
    "CreatedAt": "2025-03-11T21:46:43.736069+05:30",
    "UpdatedAt": "2025-03-11T21:46:43.736069+05:30",
    "DeletedAt": null,
    "body": "",
    "user_id": 1
  }
]
```

#### 3. Create Post (Authenticated)

**Endpoint**: `POST /posts`

**Request**:
```bash
curl -X POST http://localhost:8081/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "My First Post",
    "content": "This is the content of my first post."
  }'
```

**Response**:
```json
{
  "ID": 2,
  "CreatedAt": "2025-03-11T22:15:40.811152+05:30",
  "UpdatedAt": "2025-03-11T22:15:40.811152+05:30",
  "DeletedAt": null,
  "body": "",
  "user_id": 0
}
```

## Testing Flow

1. Register a new user
2. Sign in to get a JWT token
3. Use the JWT token to create a post
4. Retrieve all posts to verify your post was created

## Stopping the Services

To stop a service running on a specific port:

```bash
# For user service (port 8080)
lsof -i :8080 | awk 'NR>1 {print $2}' | xargs kill -9

# For post service (port 8081)
lsof -i :8081 | awk 'NR>1 {print $2}' | xargs kill -9
```

## Logging

Both services implement comprehensive logging:

1. Echo's built-in logger middleware logs all HTTP requests with details like:
   - Request method and path
   - Status code
   - Response time
   - Request size

2. Custom logger middleware adds additional logging with:
   - Request method and path
   - Client IP
   - User agent
   - Status code
   - Latency

## Project Structure

```
se_backend/
├── config.json                # Shared configuration
├── user_service/              # User authentication service
│   ├── config/                # Configuration handling
│   ├── controllers/           # Request handlers
│   ├── middlewares/           # HTTP middlewares
│   ├── models/                # Data models
│   ├── routes/                # Route definitions
│   ├── utils/                 # Utility functions
│   ├── main.go                # Entry point
│   └── go.mod                 # Go module file
└── post_service/              # Post management service
    ├── config/                # Configuration handling
    ├── controllers/           # Request handlers
    ├── middlewares/           # HTTP middlewares
    ├── models/                # Data models
    ├── routes/                # Route definitions
    ├── utils/                 # Utility functions
    ├── main.go                # Entry point
    └── go.mod                 # Go module file
``` 