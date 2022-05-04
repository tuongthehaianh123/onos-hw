ARG ONOS_BUILD_VERSION=undefined

FROM onosproject/golang-build:v1.0 as build

ENV GO111MODULE=on
ARG ONOS_MAKE_TARGET=build

COPY Makefile go.mod go.sum /go/src/github.com/tuongthehaianh123/onos-hw/
COPY cmd/ /go/src/github.com/tuongthehaianh123/onos-hw/cmd/
COPY pkg/ /go/src/github.com/tuongthehaianh123/onos-hw/pkg/
COPY vendor/ /go/src/github.com/tuongthehaianh123/onos-hw/vendor/
COPY api/ /go/src/github.com/tuongthehaianh123/onos-hw/api
COPY build/build-tools /go/src/github.com/tuongthehaianh123/onos-hw/build/build-tools

RUN cd /go/src/github.com/tuongthehaianh123/onos-hw && GOFLAGS=-mod=vendor make ${ONOS_MAKE_TARGET}

FROM alpine:3.11
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/tuongthehaianh123/onos-hw/build/_output/onos-hw /usr/local/bin/onos-hw

ENTRYPOINT ["onos-hw"]
