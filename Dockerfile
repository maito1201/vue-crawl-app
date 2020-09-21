FROM golang:1.12-alpine as builder
WORKDIR ./
RUN apk add --no-cache git
ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=
COPY ./go/go.mod ./go/go.sum ./
RUN go mod download
COPY . .
RUN env GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o app

FROM node:12.7.0-alpine
WORKDIR /myapp
COPY --from=builder /go/app ./go/app
COPY package.json ./
COPY yarn.lock ./
RUN yarn install && yarn global add @vue/cli
# install choromedriver
RUN apk add --update \
    wget \
    udev \
    ttf-freefont \
    chromium \
    chromium-chromedriver \
