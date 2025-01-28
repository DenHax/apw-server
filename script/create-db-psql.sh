#!/usr/bin/env bash

echo "Wait postgres"
sleep 20
docker cp ./sql apw-psql-serv:/sql
docker exec -it apw-psql-serv psql -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -f /sql/init.sql
docker exec -it apw-psql-serv psql -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -f /sql/data.sql
