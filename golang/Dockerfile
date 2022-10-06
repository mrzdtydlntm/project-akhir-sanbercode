# Phase 1: Build

FROM golang:1.17-alpine as build

LABEL maintainer="Mirza"

LABEL version="v1"

RUN apk update

WORKDIR /var/www

COPY ./ ./

RUN go mod download && go mod tidy

RUN go build -o server.sh

# Phase 2: Manipulation

# Image dibawah hanya sebagai runner
FROM alpine:latest

WORKDIR /var/www

ENV SERVER_PORT=8000

ENV PGHOST=database

ENV PGPORT=5432

ENV PGUSER=postgres

ENV PGPASSWORD=sanber

ENV PGDBNAME=sanber

COPY --from=build /var/www/server.sh .

CMD [ "./server.sh" ]