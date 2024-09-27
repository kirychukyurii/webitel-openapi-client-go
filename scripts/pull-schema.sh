#!/usr/bin/env bash

# https://stackoverflow.com/a/246128
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Pull the schema (stable commit)
SCHEMA="$(curl -s -L https://raw.githubusercontent.com/webitel/protos/refs/heads/main/swagger/api.json)"

# Write the schema to a file
echo "${SCHEMA}" > "${SCRIPT_DIR}/schema.json"