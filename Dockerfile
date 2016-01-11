FROM golang
ADD build-deps-linux.sh /go/src/github.com/jphastings/credence/
WORKDIR /go/src/github.com/jphastings/credence
RUN apt-get update && ./build-deps-linux.sh

ADD . /go/src/github.com/jphastings/credence
RUN make bootstrap

EXPOSE 80
EXPOSE 2733

ENTRYPOINT serverconfig/start
