# 🏥 Clinic Portal - Golang Backend Project

A secure and scalable backend application for managing patient records with two user roles: **Receptionist** and **Doctor**.

---

## 📚 Table of Contents

* [Features](#features)
* [Tech Stack](#tech-stack)
* [System Architecture](#system-architecture)
* [Setup Instructions](#setup-instructions)
* [API Documentation (Swagger)](#api-documentation-swagger)
* [Docker Support](#docker-support)
* [Directory Structure](#directory-structure)
* [Future Improvements](#future-improvements)

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

## 🧩 System Architecture

```mermaid
graph TD
    A[Client (Postman / curl)] -->|REST API| B[Golang Backend]
    B --> C[PostgreSQL DB]
    B --> D[Swagger UI]
```

---

## ⚙️ Setup Instructions

### 🔧 Backend

1. **Clone the repo**:

   ```bash
   git clone https://github.com/your-repo/clinic-portal
   cd clinic-portal
   ```

2. **Configure `.env` file**:

   ```env
   DB_URL=postgres://user:password@localhost:5432/clinic_db
   JWT_SECRET=supersecretkey
   ```

3. **Run database** (if using Docker):

   ```bash
   docker-compose up db
   ```

4. **Run backend locally**:

   ```bash
   go run cmd/main.go
   ```

---

## 🧪 API Documentation (Swagger)

1. **Install swag CLI**:

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. **Generate docs**:

   ```bash
   swag init -g cmd/main.go
   ```

3. **Visit Swagger UI**:

   ```
   http://localhost:8080/swagger/index.html
   ```

---

## 🐳 Docker Support

1. **Build and run all services**:

   ```bash
   docker-compose up --build
   ```

2. Services:

   * `backend`: Golang app with API and Swagger
   * `db`: PostgreSQL with seeded data

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
├── utils/
├── docs/                 # Swagger docs
├── go.mod / go.sum
├── .env
├── Dockerfile
└── docker-compose.yml
```

---

## 💡 Future Improvements

* Password encryption and registration
* Admin role for managing users
* Search and pagination
* Frontend interface
* Role-based dashboards

---
