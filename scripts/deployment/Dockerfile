FROM somq14/holo-back-base:0.0.1 as builder
WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./ ./
RUN go build -x -o ./build/app ./cmd/app/main.go

FROM alpine:3.11 as app
WORKDIR /work

COPY ./config /work/config
COPY --from=builder /work/build/app /usr/local/bin

CMD [ "app" ]
