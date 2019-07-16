FROM golang:1.12.7-alpine3.10 as build

RUN apk add git

COPY . /app
WORKDIR /app
RUN go install
RUN go build

FROM alpine:3.10

COPY --from=build /app/main .
CMD ["./main"]