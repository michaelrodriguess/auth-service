# 🔐 Auth Service

An authentication service built with **Go (Golang)** that provides user registration, login, token validation, and JWT-based authentication. Ideal for applications that need a decoupled authentication module.

## ✨ Features

- ✅ User registration (`/register`)
- ✅ User login (`/login`)
- ✅ Authenticated user info (`/me`)
- ✅ JWT generation and validation
- ✅ Secure credential hashing
- ✅ Middleware for protected routes
- ✅ Docker and Makefile for easy setup

---

## 🛍 Available Routes

### `POST /register`
Registers a new user.

**Example Body:**
```json
{
  "email": "email@domain.com",
  "password": "123456",
  "role": "admin" or "user"
}
```

**Response:**
```json
{
  "token": "0CT0LHOIvqzmRTY3aTU5-kkvkeukFiFjoqG5N9FOcFM......",
  "email": "email@domain.com",
  "role": "admin" or "user"
}
```

---

### `POST /login`
Authenticates a user and returns a JWT.

**Example Body:**
```json
{
  "email": "candango@dev.com",
  "password": "123456"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

---

### `GET /me`
Returns authenticated user data.

**Required Header:**
```http
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "id",
  "email": "email@domain.com",
  "role": "admin" or "user",
  "created_at": timestamp,
  "updated_at": timestamp
}
```

---

## Installation

### Requirements

- Go 1.18 or higher
- Docker
- Make (optional, for ease of use)


## 🐳 Running MongoDB on Docker

```bash
docker-compose up --build -d
```

---

## 🛠️ Makefile Commands

Useful commands for development:

```bash
make run       # Run the app locally
make build     # Build the project
make test      # Run tests (when added)
make lint      # Run linter (when configured)
```

## Running the Go Server

```bash
make run
```

---

## ⚙️ Environment Variables

Create a `.env` file with the following content:

```env
PORT=8080
DB_URI=mongodb://mongo:27017/authdb
JWT_SECRET=your_secure_secret_key
```

---

## 📌 Future Roadmap

- [ ] Password reset flow
- [ ] Refresh token support
- [ ] RBAC (Role-based access control)

---

## Developed by

Michael Araujo Rodrigues — [@michaelrodriguess](https://github.com/michaelrodriguess)

