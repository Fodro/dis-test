#!/usr/bin/env sh

set -e

COLOR="\e[34m"
DEFAULT_COLOR="\e[0m"

printf "$COLOR\n### %s$DEFAULT_COLOR\n" "Migrating database ..."
make migrate-up
