FROM golang:alpine

RUN apk add git

COPY . /app
WORKDIR /app
RUN go install
RUN go build
ENTRYPOINT ./main