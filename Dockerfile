# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN mkdir /src
ADD . /src
WORKDIR /src

ENV db_user="root"
ENV db_pass=
ENV db_host=host.docker.internal
ENV db_name="siremun-dokumen"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build /src

EXPOSE 8080

CMD ["./go-rest-api"]