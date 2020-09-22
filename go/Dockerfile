FROM golang:1.12-alpine as builder
WORKDIR ./
RUN apk add --no-cache git
ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=
COPY ./go.mod ./go.sum ./main.go ./
RUN go mod download
RUN env GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o /go-api

FROM node:12.7.0-alpine
WORKDIR /myapp
COPY --from=builder /go-api .
# install choromedriver
RUN apk add --update \
    wget \
    udev \
    ttf-freefont \
    chromium \
    chromium-chromedriver \
