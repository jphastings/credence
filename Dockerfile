FROM jphastings/credence:dependencies
ADD . /go/src/github.com/jphastings/credence
WORKDIR /go/src/github.com/jphastings/credence
RUN make bootstrap

EXPOSE 80
EXPOSE 2733

ENTRYPOINT serverconfig/start
