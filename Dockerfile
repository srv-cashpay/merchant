FROM golang:1.24.4

RUN mkdir /app

WORKDIR /app
COPY ./configs/firebase-service-account.json /app/configs/firebase-service-account.json

ADD go.mod .
ADD go.sum .

RUN go mod download