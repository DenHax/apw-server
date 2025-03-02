#!/usr/bin/env bash

psql -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -f /sql/init.sql
psql -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" -f /sql/data.sql
