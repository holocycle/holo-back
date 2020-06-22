FROM golang:1.14-alpine as base
WORKDIR /work

RUN apk add --no-cache bash curl git

ENV CGO_ENABLED 0
RUN GO111MODULE=on  go get -x github.com/golangci/golangci-lint/cmd/golangci-lint@v1.26.0 \
  && GO111MODULE=off go get -x github.com/pressly/goose/cmd/goose \
  && GO111MODULE=off go get -x github.com/oxequa/realize \
  && cp $(go env GOPATH)/bin/* /usr/local/bin \
  && rm -r $(go env GOPATH) $(go env GOCACHE)
