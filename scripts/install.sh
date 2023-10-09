#!/usr/bin/env sh

if ! [ -x "$(command -v docker-compose)" ]; then
  echo 'Error: docker-compose is not installed.' >&2
  exit 1
fi

set -e

COLOR="\e[34m"
DEFAULT_COLOR="\e[0m"

printf "$COLOR\n### %s$DEFAULT_COLOR\n" "Creating docker-compose .env ..."
if [ ! -f "$ENV_FILE" ]; then
  cp "$ENV_FILE.dist" "$ENV_FILE"
  echo "Successfully created"
else
  echo "Docker-compose .env file already exists"
fi

set -a
# shellcheck source=deployments/.env
[ -f "$ENV_FILE" ] && . "$ENV_FILE"
set +a

printf "$COLOR\n### %s$DEFAULT_COLOR\n" "Building containers ..."
if [ -z "$CI" ]; then
  if [ "$(uname)" = "Linux" ]; then
    USER_ID=$(id -u)
  else
    USER_ID=1000
  fi

  ${DOCKER_COMPOSE} build --build-arg UID=${USER_ID}
else
  ${DOCKER_COMPOSE} build
fi
