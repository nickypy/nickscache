FROM golang:1.12.7-alpine3.10 as build

RUN apk add --update --no-cache ca-certificates git

RUN mkdir /out
WORKDIR /out
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

COPY . .
RUN go install
RUN go build -o /out/stache

FROM alpine:3.10

COPY --from=build /out/stache .
ENTRYPOINT ["/stache"]