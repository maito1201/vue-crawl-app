FROM node:12.7.0-alpine

WORKDIR /myapp

COPY package.json ./
COPY yarn.lock ./

RUN yarn install && yarn global add @vue/cli
# install golang
RUN apk add --update --no-cache vim git make musl-dev go curl && \
    export GOPATH= && \
    export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH && \
    export GOBIN=$GOROOT/bin && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin && \
    export GO111MODULE=on
# install choromedriver
RUN apk add --update \
    wget \
    udev \
    ttf-freefont \
    chromium \
    chromium-chromedriver \