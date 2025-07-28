# 🔐 Go JWT Login Application

A simple JWT-based authentication API built with Golang. This project uses JSON Web Tokens (JWT) for secure user login and protected route access. It is containerized using Docker and uses MySQL as its database.

## 📁 Project Structure

```
|   .env
|   .gitignore
|   docker-compose.yaml
|   Dockerfile
|   go-jwt.exe
|   go-jwt.exe~
|   go.mod
|   go.sum
|   main.go
|   README.md
|
+---configuration
|   +---initializers
|   |       dbConnection.go
|   |       loadEnvViariables.go
|   |       syncDatabase.go
|   |
|   +---loggerHandler
|   |       loggerHandler.go
|   |
|   +---middleware
|   |       requireAuth.go
|   |
|   \---validation
|           validateUser.go
|
+---controllers
|   |   login.go
|   |   logout.go
|   |   signup.go
|   |   userController.go
|   |
|   +---requestDTO
|   |       loginDTO.go
|   |       signUpDTO.go
|   |
|   \---routes
|           routes.go
|
+---domain
|   \---user
|       |   userDomain.go
|       |
|       \---services
|               createUser.go
|               findUser.go
|               userInterface.go
|
+---dtos
|       userDto.go
|
+---errorHandler
|       errorHandler.go
|
\---models
    \---user
        +---entities
        |       userModel.go
        |
        \---repositories
                createUserRepository.go
                findUserRepository.go
                userRepository.go

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