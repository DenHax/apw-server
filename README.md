# APW server

Backend in Golang: Gin, SQLX, PQ, Cleanenv

DB connect: in .env or .env-compose

API URL example: http://localhost:8080/api/

Database: [repository](https://github.com/DenHax/apw-db)

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

Docker compose manual:

```sh
make compose-create-env # or modify .env-example for test compose config
ENV=compose . ./scripts/activate-env.sh # activate environment variables
make compose-run # wait 10-15 second before run next script or remove -d from script for view postgres' logs
make create-db # or remove cp data.sql from script for clean init database model
```

Manual without storage:

```go
go mod tidy
CONFIG=./configs/config.yaml go run cmd/apw/main.go
```
