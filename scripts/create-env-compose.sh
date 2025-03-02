#!/usr/bin/env bash

if [ -f ./.env-compose ]; then
  echo "Error: .env-compose file already exists"
else
  cat <<EOL >.env-compose
POSTGRES_HOST=storage
POSTGRES_PORT=5432
POSTGRES_USER=admin
POSTGRES_DB=apw
POSTGRES_PASSWORD=p4ssw0rd
SSL_MODE=disable
APP_PORT=8080
POSTGRES_URL=postgres://admin:p4ssword@storage:5432/apw?sslmode=disable
EOL
  echo ".env-compose file created successfully"
fi
