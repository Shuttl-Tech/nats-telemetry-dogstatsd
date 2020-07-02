FROM golang:1.14 as build
ENV CGO_ENABLED=0

ADD . /src
WORKDIR /src

RUN make build


FROM alpine:latest as certs
RUN apk update && apk add ca-certificates


FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/_build/nats-telemetry-dogstatsd /nats-telemetry-dogstatsd

ENTRYPOINT ["/nats-telemetry-dogstatsd"]