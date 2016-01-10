FROM golang
ADD . /go/src/github.com/jphastings/credence
WORKDIR /go/src/github.com/jphastings/credence
RUN build-deps-linux.sh
RUN make bootstrap

ENTRYPOINT serverconfig/start

EXPOSE 80
EXPOSE 2733