FROM golang:1.19.2-alpine3.16


LABEL maintainer="raudra.nitt@gmail.com"

RUN apk add build-base bash git


WORKDIR /profile-service

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /profile-service

RUN go mod tidy

RUN go build -o /usr/local/bin/profile-service

CMD profile-service
