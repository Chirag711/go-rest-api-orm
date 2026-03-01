# GORM REST API with Go, Gin and MySQL

A simple and clean **RESTful API built with Golang using Gin Framework and GORM ORM** connected to a **MySQL database**.
This project demonstrates how to build a basic **CRUD API with proper project structure**.

---

## Features

* REST API using **Gin Framework**
* **GORM ORM** for database operations
* **MySQL database integration**
* Full **CRUD operations**
* Clean **MVC style folder structure**
* Auto database migration
* JSON request and response handling

---

## Tech Stack

* **Golang**
* **Gin Web Framework**
* **GORM ORM**
* **MySQL**
* **REST API**

---

## Project Structure

```
go-rest-api-orm
│
├── main.go
│
├── config
│   └── db.go
│
├── models
│   └── user.go
│
├── controllers
│   └── userController.go
│
└── routes
    └── routes.go
```

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/Chirag711/go-rest-api-orm.git
cd go-rest-api-orm
```

---

### Initialize Go Module

```bash
go mod init go-rest-api-orm
```

---

### Install Dependencies

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

---

## Database Setup

Create a MySQL database:

```sql
CREATE DATABASE gormapi;
```

---

## Configure Database Connection

Edit **config/db.go**

```go
dsn := "root:password@tcp(127.0.0.1:3306)/gormapi?charset=utf8mb4&parseTime=True&loc=Local"
```

Replace:

```
root
password
```

with your MySQL credentials.

---

## Run the Application

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## API Endpoints

| Method | Endpoint       | Description     |
| ------ | -------------- | --------------- |
| POST   | /api/users     | Create user     |
| GET    | /api/users     | Get all users   |
| GET    | /api/users/:id | Get single user |
| PUT    | /api/users/:id | Update user     |
| DELETE | /api/users/:id | Delete user     |

---

## Example Request

### Create User

**POST**

```
/api/users
```

Request Body

```json
{
  "name": "Chirag",
  "email": "chirag@gmail.com",
  "age": 25
}
```

---

## Example Response

```json
{
  "id": 1,
  "name": "Chirag",
  "email": "chirag@gmail.com",
  "age": 25
}
```

---

## Auto Migration

GORM automatically creates the database table using:

```go
config.DB.AutoMigrate(&models.User{})
```

Table created:

```
users
```

Fields:

```
id
created_at
updated_at
deleted_at
name
email
age
```

---

## Future Improvements

* JWT Authentication
* Request validation
* Pagination
* Environment variables (.env)
* Docker support
* Logging middleware
* Clean architecture

---

## Contributing

Contributions are welcome. Feel free to fork the repository and submit a pull request.

---

## License

This project is open source and available under the **MIT License**.

---

## Author

Developed using **Golang + Gin + GORM + MySQL**.
