# 🔐 Go JWT Login Application

A simple JWT-based authentication API built with Golang. This project uses JSON Web Tokens (JWT) for secure user login and protected route access. It is containerized using Docker and uses MySQL as its database.

## 📁 Project Structure

```
|   .env
|   docker-compose.yaml
|   Dockerfile
|   go-jwt.exe
|   go-jwt.exe~
|   go.mod
|   go.sum
|   main.go
|   
+---controllers
|       userController.go
|       
+---initializers
|       dbConnection.go
|       loadEnvViariables.go
|       syncDatabase.go
|
+---middleware
|       requireAuth.go
|
\---models
        userModel.go
```

---

## 🚀 Features

- User registration and login
- Password hashing using bcrypt
- JWT token generation on login
- JWT validation middleware for protected routes
- MySQL database integration with GORM
- .env-based configuration loading
- Dockerized application with MySQL

---

## 🔧 Prerequisites

- Go 1.20+
- Docker & Docker Compose

---