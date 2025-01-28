#!/usr/bin/env bash

docker run --name=apw-psql-serv \
  -e POSTGRES_USER="$POSTGRES_USER" \
  -e POSTGRES_PASSWORD="$POSTGRES_PASSWORD" \
  -e POSTGRES_DB=POSTGRES_DB \
  -p "$POSTGRES_PORT":5432 \
  postgres:16-alpine
