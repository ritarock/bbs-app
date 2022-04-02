FROM golang:1.18.0-buster AS builder

WORKDIR /app
COPY . /app
RUN make build

FROM golang:1.18.0-buster
COPY --from=builder /go/bin/backend /go/bin/backend
EXPOSE 8080
CMD [ "/go/bin/backend" ]
