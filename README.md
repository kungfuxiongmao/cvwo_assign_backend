# Sample Go App

A backend web application built with **Go**, **Gin**, **GORM**, **Bcrypt** and **PostgreSQL**.  
This project provides APIs for managing users, posts, topics, and comments.


## Features

- RESTful API using **Gin**
- PostgreSQL database with **GORM** ORM
- User authentication using **Bcrypt** and middleware support
- Modular handlers for topics, posts, comments, and users

## Prerequisites

- Go (v1.18+)
- PostgreSQL
- Git

## Installation and Setup
### 1. Clone the Repository
```bash
git clone https://github.com/kungfuxiongmao/cvwo_assign_backend.git
cd cvwo_assign_backend
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Create Database
### 4. Configure environment Variables
Set up the following parameters in a .env file
```bash
# Database variables
DB_HOST=
DB_PORT=
DB_USER=
DB_PASS=
DB_NAME=
DB_SSL=

# Frontend port (CORS Settings)
FRONTEND=

# Cookies & Auth Setup
JWT_KEY=
GO_ENV="local"/"production"
```
### 5. Run the Application
```bash
go run main.go
```

## AI Declaration
AI has been used in implementing the project, in explaining concepts and aiding in decision making, as well as in code review.

