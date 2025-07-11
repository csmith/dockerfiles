# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r2","apk:libmagic":"5.46-r2","apk:make":"4.4.1-r3","apk:musl":"1.2.5-r10","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b"}

FROM reg.c5h.io/golang@sha256:d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r3 \
        musl=1.2.5-r10 \
        file=5.46-r2 \
        libmagic=5.46-r2 \
        musl=1.2.5-r10 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
