FROM golang:1.17-alpine

WORKDIR /go/src

COPY ../ .

RUN go install github.com/cosmtrek/air@v1.27.3

WORKDIR cmd

CMD ["air", "-c", ".air.toml"]