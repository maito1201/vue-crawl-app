FROM node:12.7.0-alpine

WORKDIR /myapp

COPY package.json ./
COPY yarn.lock ./

RUN yarn install && yarn global add @vue/cli

