FROM golang:1.19.2-alpine3.16

RUN apk add --update alpine-sdk

WORKDIR /app

COPY . /app

RUN make

EXPOSE 16761

ENTRYPOINT ["build/apiserver"]