#!/bin/bash

FROM alpine as firebase-emulator

WORKDIR /firebase-emulator
COPY . .

RUN apk update && apk add --no-cache curl bash npm openjdk17-jre
RUN echo "Installing Firebase CLI" \  
    && npm install -g firebase-tools \
    && firebase --version \
    && echo "Firebase CLI installed"

RUN echo "Installing Firebase Functions Emulator - firestore, ui, storage" \
    && firebase setup:emulators:firestore \
    && firebase setup:emulators:ui \ 
    && firebase setup:emulators:storage \
    && echo "Firebase Functions Emulator installed. Ready to run."
