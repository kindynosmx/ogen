FROM golang:1.16beta1-buster

WORKDIR /build

RUN git clone https://github.com/kindynosmx/ogen

WORKDIR /build/ogen

RUN go build cmd/ogen/ogen.go

ENTRYPOINT ["./scripts/full-node/run.sh"]




