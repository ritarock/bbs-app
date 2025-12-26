FROM node:25-alpine

RUN npm install -g @typespec/compiler

WORKDIR /app

COPY . /app
