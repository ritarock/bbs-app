FROM golang:1.17.3-buster AS builder

WORKDIR /work

COPY . /work
RUN cd cmd/bbs-app/ && go install

FROM golang:1.17.3-buster
COPY --from=builder /go/bin/bbs-app /go/bin/bbs-app
EXPOSE 8080
CMD ["/go/bin/bbs-app"]
