#!/bin/bash

# Pull the latest changes from the repository
git pull origin main
git submodule update --recursive

# Remove unused docker images
docker image prune -af

# Build and run the docker containers
docker compose build
docker compose up -d --force-recreate
