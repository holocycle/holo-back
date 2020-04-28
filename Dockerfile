FROM golang:1.14-alpine as builder
WORKDIR /work

RUN apk add git \
  && go get -x github.com/oxequa/realize

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./ ./
RUN go build main.go

FROM alpine:3.11 as app
COPY --from=builder /work/main /usr/local/bin
CMD [ "main" ]
