FROM alpine

COPY gopath/bin/lamehellogo /go/bin/lamehellogo

ENTRYPOINT /go/bin/lamehellogo