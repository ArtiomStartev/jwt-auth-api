# JWT Authentication in Go

## Introduction
Hey there 👋! This repository demonstrates how to implement JWT Authentication in Golang REST APIs using the Fiber Web Framework, PostgreSQL database, and GORM. This tutorial guides you through creating a robust API with user authentication, covering functionalities like registering, logging in, fetching user details, and logging out.

---

## What is JWT?
JSON Web Token (JWT) is an open standard (RFC 7519) for securely transmitting information between parties as a JSON object. This information is verified and trusted using a digital signature, ensuring the token’s legitimacy.

### Key Features of JWT:
- **Authentication**: JWTs are commonly used for user authentication.
- **Claims**: Tokens assert claims like "logged in as admin."
- **Client-Server Communication**: JWTs enable the server to trust the client’s requests.

---

## Features of the Project
- **Register**: Create a new user.
- **Login**: Authenticate an existing user and generate a JWT.
- **Get Active User**: Retrieve details of the currently logged-in user.
- **Logout**: Log out the user by expiring their token.

---

## Prerequisites
Ensure the following tools are installed:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Go-Fiber Framework](https://gofiber.io/)
- [GORM ORM](https://gorm.io/)

---

## Installation

### Step 1: Clone the Repository
```bash
git clone https://github.com/ArtiomStartev/jwt-auth-api.git
cd jwt-auth-api
```

### Step 2: Install Dependencies
```bash
go get -u github.com/gofiber/fiber/v2
```
```bash
go get -u gorm.io/gorm
```
```bash
go get -u gorm.io/driver/postgres
```
```bash
go get github.com/golang-jwt/jwt
```

---

## Setup

### Database Configuration
Update the `database/database.go` file with your PostgreSQL credentials:
```go
const (
    host     = "localhost"
    port     = 5432
    user     = "your_username"
    password = "your_password"
    dbname   = "jwt-auth-api"
)
```

Run the database migrations:
```bash
go run main.go
```

---

## Project Structure
```
.
├── controller/         # Contains API route handlers
├── database/           # Database connection and migrations
├── models/             # Database models (e.g., User)
├── routes/             # Route setup
├── main.go             # Entry point of the application
└── README.md           # Documentation
```

---

## API Endpoints

### 1. **Register User**
- **Endpoint**: `/user/register`
- **Method**: POST
- **Request Body**:
  ```json
  {
      "name": "John Doe",
      "email": "john@example.com",
      "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
      "data": {
          "ID": 1,
          "name": "John Doe",
          "email": "john@example.com"
      },
      "error": null
  }
  ```

### 2. **Login User**
- **Endpoint**: `/user/login`
- **Method**: POST
- **Request Body**:
  ```json
  {
      "email": "john@example.com",
      "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
      "data": {
          "email": "john@example.com"
      },
      "error": null
  }
  ```

### 3. **Get Active User**
- **Endpoint**: `/user/get-user`
- **Method**: GET
- **Headers**:
  ```http
  Cookie: jwt=your_jwt_token
  ```
- **Response**:
  ```json
  {
      "data": {
          "ID": 1,
          "name": "John Doe",
          "email": "john@example.com"
      },
      "error": null
  }
  ```

### 4. **Logout User**
- **Endpoint**: `/user/logout`
- **Method**: POST
- **Response**:
  ```json
  {
      "data": null,
      "error": null
  }
  ```

---

## Running the Application
Start the server:
```bash
go run main.go
```

The application will run on `http://localhost:8000`.

---

## Technologies Used
- **Golang**: Programming language
- **Fiber**: Web framework
- **PostgreSQL**: Database
- **GORM**: ORM for Golang
- **JWT**: Authentication mechanism

---

Happy coding 🚀!