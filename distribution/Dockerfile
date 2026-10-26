# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.47-r2","apk:libmagic":"5.47-r2","apk:make":"4.4.1-r4","apk:musl":"1.2.6-r3","image:base":"091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47","image:golang":"c04403d382d3f01c5081af9cf8b27648cf75d06c3111ca5310304a8b3e8e454f"}

FROM reg.c5h.io/golang@sha256:c04403d382d3f01c5081af9cf8b27648cf75d06c3111ca5310304a8b3e8e454f AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r4 \
        musl=1.2.6-r3 \
        file=5.47-r2 \
        libmagic=5.47-r2 \
        musl=1.2.6-r3 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
