FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o go-app-rest-api

CMD ["./go-app-rest-api"]

ARG PORT=8080
ENV PORT=$PORT
EXPOSE $PORT

ENV NAME=go-app-rest-api