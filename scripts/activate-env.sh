#!/usr/bin/env bash

if [ -n "${ENV}" ]; then
  filenv="./.env-${ENV}"
else
  filenv="./.env"
fi

source "${filenv}"
