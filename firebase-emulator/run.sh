#!/bin/bash

CURR_DIR=$(dirname $0)
IMAGE_NAME='firebase-emulator:latest'

if docker image inspect $IMAGE_NAME >/dev/null 2>&1; then
    echo "Image exists locally"
else
    echo "Image does not exist locally"
    docker build -t $IMAGE_NAME -f $CURR_DIR/Dockerfile .
fi

docker run --rm \
  -p 9199:9199 \
  -p 9099:9099 \
  -p 9080:9080 \
  -p 4000:4000 \
  --name firebase-emulator \
  -v $CURR_DIR:/firebase-emulator \
  $IMAGE_NAME \
  firebase emulators:start --project fluentify-test
