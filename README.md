# APW server

Backend in Golang: Gin, SQLX, PQ, Cleanenv

DB connect: in .env

API URL example: http://localhost:8080

# Handlers

- Employee:

  - /employee/ - GET all employees that works as 'оператор реактора'
  - /employee/:id - UPDATE employee by ID (employee_id)

- Subsystem:

  - /subsystem/ - GET all subssystems that named as 'реактор'

- Fuel road:

  - /fuel-road/ - GET all fuel roads

- Upload:
  - /upload/ - GET all uploads
  - /upload/:id - DELETE upload by id (load_date)
  - /upload/- POST upload with upload's data from anywhere
  - /upload/report/ GET upload's report data

# Architecture:

React frontend -> HTTP -> Handler -> Service -> Repository -> PostgreSQL

# Get start

```go
go mod tidy
go run cmd/apw/main.go
```
