#!/bin/bash

FROM alpine as firebase-emulator

WORKDIR /firebase-emulator

RUN apk update && apk add --no-cache curl bash npm openjdk17-jre
RUN echo "Installing Firebase CLI" \
    && npm install -g firebase-tools \
    && firebase --version \
    && echo "Firebase CLI installed"
