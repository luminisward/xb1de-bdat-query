FROM golang:1.22 as builder

WORKDIR /app

ARG GO111MODULE
ARG GOPROXY

ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .

FROM alpine:3.19

COPY --from=builder /app/main /main

ENV GIN_MODE=release

CMD /main