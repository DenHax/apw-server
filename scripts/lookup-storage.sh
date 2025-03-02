#!/usr/bin/env bash

PSQL_DOCKER_ID=$(docker ps --filter='name=apw-psql' --format='{{.ID}}')

if [ -z "$PSQL_DOCKER_ID" ]; then
  echo "Container for postgres not found. Script stopped."
  exit 1
fi

docker exec -it "$PSQL_DOCKER_ID" /usr/bin/env bash -c "psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -p ${POSTGRES_PORT}"
