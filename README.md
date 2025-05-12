# Divine Pizzeria Backend

A modular, RESTful API written in Go for managing the backend operations of a fictional pizzeria. Built with the `gin-gonic` framework and PostgreSQL, structured for scalability and maintainability using a layered architecture.

---

## 🧱 Tech Stack

- **Language**: Go 1.24.2
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Environment Config**: `godotenv`
- **Containerization**: Docker & Docker Compose

---

## 📁 Project Structure
```
divine-pizzeria-backend/
│
├── config/ # App and DB configuration
│ ├── config.go
│ └── database.go
│
├── example_bash/ # Sample migration scripts
│
├── internal/
│ └── database/
│   └── migrations/ # .sql files for DB schema
│
├── modules/ # Versioned feature modules
│ └── module_name/
│   └── v1/
│       ├── handlers/
│       ├── models/
│       ├── repositories/
│       └── routes/
│
├── router/ # Routing setup
│ ├── router.go
│ └── routes.go
│
├── .env.example # Template for environment variables
├── docker-compose.yml # Sets up PostgreSQL container
├── go.mod # Module definition and dependencies
├── go.sum
└── main.go # Application entry point
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/divine-pizzeria-backend.git
cd divine-pizzeria-backend
```

### 2. Setup Environment
Create a .env file from the example provided:

```bash
cp .env.example .env
```
### 3. Run PostgreSQL with Docker Compose
Make sure Docker is installed and running:

```bash
docker-compose up -d
```
### 4. Apply Migrations
If you want to apply SQL migrations, use the scripts in example_bash/RUN-MIGRATION.bash

### 5. Run the Application
```bash
go run main.go
```
By default, the API will run on http://localhost:8080.

## 🧪 API Versioning
Each module is organized by version inside the modules/ folder. For example:

```
GET /api/v1/pizzas
POST /api/v1/orders
```
## 🛠 Available Scripts
You can find useful scripts (like DB initialization and reset) inside the example_bash/ folder.

## 📦 Dependencies
Key dependencies listed in go.mod:

- github.com/gin-gonic/gin

- github.com/joho/godotenv

- github.com/lib/pq

- gorm.io/gorm

- gorm.io/driver/postgres

Run the following if needed:

```bash
go mod tidy
```
## 🧩 Environment Variables
The application uses environment variables for configuration. Create a `.env` file based on the `.env.example` template.

## 🐳 Docker
You can easily extend docker-compose.yml to also run the Go API in a container. Currently, it's used for the database only.

## 📚 License
This project is licensed under the MIT License.
