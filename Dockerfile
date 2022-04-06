# syntax=docker/dockerfile:1

FROM golang:1.18.0-alpine

WORKDIR /

COPY . ./
RUN go mod download

RUN go build -o /app

ENTRYPOINT /app

EXPOSE 3000

CMD [ "/app" ]
