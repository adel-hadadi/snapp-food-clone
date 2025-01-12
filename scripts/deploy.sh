#!/bin/bash

ENVIRONMENT=$1

if [ "$ENVIRONMENT" == "production" ]; then
  DOCKERFILE="app.prod.Dockerfile"
elif [ "$ENVIRONMENT" == "development" ]; then
  DOCKERFILE="app.dev.Dockerfile"
else
  echo "Usage: ./deploy.sh [production|development]"
fi

APP_ENV=$ENVIRONMENT DOCKERFILE=$DOCKERFILE docker compose up --build
