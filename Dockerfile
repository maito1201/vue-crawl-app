FROM node:12.7.0-alpine

WORKDIR /myapp

COPY package.json ./
COPY yarn.lock ./

RUN yarn install && yarn global add @vue/cli
RUN apk add --update --no-cache vim git make musl-dev go curl && \
    export GOPATH= && \
    export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH && \
    export GOBIN=$GOROOT/bin && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin && \
    export GO111MODULE=on
RUN curl -O https://chromedriver.storage.googleapis.com/77.0.3865.40/chromedriver_linux64.zip && \
    unzip chromedriver_linux64.zip && mv chromedriver /usr/local/bin/
