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

COPY --from=build /var/www/server.sh .

CMD [ "./server.sh" ]