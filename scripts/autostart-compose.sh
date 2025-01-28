#!/usr/bin/env bash

if [ -f ./.env-compose ]; then
  echo "Error: .env file already exists"
else
  cat <<EOL >.env-compose
POSTGRES_HOST=storage
POSTGRES_PORT=5432
POSTGRES_USER=admin
POSTGRES_DB=apw
POSTGRES_PASSWORD=p4ssw0rd
SSL_MODE=disable
APP_PORT=8080
POSTGRES_URL=postgres://\${POSTGRES_USER}:\${POSTGRES_PASSWORD}@\${POSTGRES_HOST}:\${POSTGRES_PORT}/\${POSTGRES_DB}?sslmode=\${SSL_MODE}
EOL
  echo ".env-compose file created successfully"
fi

if [ -f ./.env-compose ]; then
  eval "$(grep -v '^#' ./.env-compose | xargs -d '\n' -I {} echo export {})"

  if [ $? -eq 0 ]; then
    echo "Environment activation: succeeded"
  else
    echo "Error: Failed to export environment variables"
  fi
else
  echo "Error: .env file not found"
fi
