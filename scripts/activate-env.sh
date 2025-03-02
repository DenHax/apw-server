#!/usr/bin/env bash

if [ -n "${ENV}" ]; then
  filenv="./.env-${ENV}"
else
  filenv="./.env"
fi

if [ -f "./$filenv" ]; then
  while IFS= read -r line; do
    if [[ $line != \#* ]] && [[ $line != "" ]]; then
      key=$(echo "$line" | cut -d '=' -f 1 | xargs)
      value=$(echo "$line" | cut -d '=' -f 2- | xargs)
      export "$key=$value"
    fi
  done <"./$filenv"
else
  echo "File ${filenv} don't found"
  exit 1
fi
