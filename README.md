Here's your **corrected and cleaned-up version** of the README with proper **numbering**, **consistent formatting**, and no duplicate sections:

---

# 🏥 Clinic Portal - Golang Backend Project

A secure and scalable backend application for managing patient records with two user roles: **Receptionist** and **Doctor**.

---

## 📚 Table of Contents

1. [Features](#-features)
2. [Tech Stack](#-tech-stack)
3. [System Architecture](#-system-architecture)
4. [Setup Instructions](#-setup-instructions)
5. [How to Run](#-how-to-run)
6. [API Documentation (Swagger)](#-api-documentation-swagger)
7. [Docker Support](#-docker-support)
8. [Directory Structure](#-directory-structure)
9. [Future Improvements](#-future-improvements)

---

## ✅ Features

### 🧾 Authentication

* JWT-based login system
* Role-based route protection

### 👩‍⚕️ Receptionist Portal

* Create new patients
* View patient list
* Edit and delete patients

### 👨‍⚕️ Doctor Portal

* View all patients
* Update patient diagnosis

---

## 💻 Tech Stack

| Layer     | Technology        |
| --------- | ----------------- |
| Backend   | Golang + Gin      |
| Database  | PostgreSQL + GORM |
| Docs      | Swagger (swaggo)  |
| Container | Docker            |

---


---

## ⚙️ Setup Instructions

### 1. Clone the Repo

```bash
git clone https://github.com/your-repo/clinic-portal
cd clinic-portal
```

### 2. Configure Environment Variables

Create a `.env` file (or set directly in `config.go` if hardcoded):

```env
DB_URL=postgres://user:password@localhost:5432/clinic_db
JWT_SECRET=supersecretkey
```

### 3. (Optional) Run PostgreSQL with Docker

```bash
docker-compose up db
```

---

## 🚀 How to Run

### 1. Initialize Go Module

```bash
go mod init clinic-portal
```

Creates the `go.mod` file to manage dependencies.

---

### 2. Install Dependencies

```bash
go mod tidy
```

Installs:

* `github.com/gin-gonic/gin`
* `github.com/dgrijalva/jwt-go` (for JWT auth)

---

### 3. Start the Server

```bash
go run cmd/main.go
```

✅ App runs at: [http://localhost:8080]

---

## 🧪 API Documentation (Swagger)

### 1. Install swag CLI

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Generate Swagger Docs

```bash
swag init -g cmd/main.go
```

### 3. View Swagger UI

Open in your browser:

```
http://localhost:8080/swagger/index.html
```

---

## 🐳 Docker Support

### 1. Build and Run All Services

```bash
docker-compose up --build
```

### 2. Services

* `backend`: Golang app with APIs and Swagger
* `db`: PostgreSQL database with seeded data

---

## 📁 Directory Structure

```
clinic-portal/
│
├── cmd/
│   └── main.go
├── controller/
├── middleware/
├── model/
├── routes/
├── store/
├── utils/
├── docs/                 # Swagger docs
├── go.mod / go.sum
├── .env
├── Dockerfile
└── docker-compose.yml
```

---

## 💡 Future Improvements

* Password encryption & user registration
* Admin role for managing staff
* Patient search & pagination
* Frontend dashboard for both roles
* Role-based dashboard views

---
