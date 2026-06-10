# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.47-r2","apk:libmagic":"5.47-r2","apk:make":"4.4.1-r4","apk:musl":"1.2.6-r2","image:base":"932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929","image:golang":"f6a5f3b1a6a705f2608ba10bb77e9d9e855775f50c8ccd6d159ecea6fdb8d1f6"}

FROM reg.c5h.io/golang@sha256:f6a5f3b1a6a705f2608ba10bb77e9d9e855775f50c8ccd6d159ecea6fdb8d1f6 AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r4 \
        musl=1.2.6-r2 \
        file=5.47-r2 \
        libmagic=5.47-r2 \
        musl=1.2.6-r2 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
