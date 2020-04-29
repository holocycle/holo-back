FROM golang:1.14-alpine as builder
WORKDIR /work

RUN apk add git \
  && go get -x github.com/oxequa/realize

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./ ./
RUN go build -o build/app ./cmd/app/main.go

FROM alpine:3.11 as app
WORKDIR /work

COPY ./config /work/config
COPY --from=builder /work/build/app /usr/local/bin

CMD [ "app" ]
