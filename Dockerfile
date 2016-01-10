FROM golang
RUN apt-get update
RUN echo "deb http://httpredir.debian.org/debian stretch main" >> /etc/apt/sources.list
# TODO: install protobuf v3
RUN apt-get install -y pkg-config libsodium-dev libczmq-dev

ADD . /go/src/github.com/jphastings/credence
RUN cd /go/src/github.com/jphastings/credence; make bootstrap

ENTRYPOINT /go/src/github.com/jphastings/credence/serverconfig/start

EXPOSE 80
EXPOSE 2733