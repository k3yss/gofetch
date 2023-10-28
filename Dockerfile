FROM golang:1.21.3-bookworm

RUN mkdir /build

ADD . /build/

WORKDIR /build/

RUN go build .

RUN apt-get update && apt-get install -y lsb-release

CMD ["./gofetch"]
