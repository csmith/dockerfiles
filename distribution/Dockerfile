# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r2","apk:libmagic":"5.46-r2","apk:make":"4.4.1-r2","apk:musl":"1.2.5-r8","image:base":"a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af","image:golang":"dd6035aab3b9568fd9cf56c1bfe223c945f1f7ed64f59c65230423bfb6fd136d"}

FROM reg.c5h.io/golang@sha256:dd6035aab3b9568fd9cf56c1bfe223c945f1f7ed64f59c65230423bfb6fd136d AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r2 \
        musl=1.2.5-r8 \
        file=5.46-r2 \
        libmagic=5.46-r2 \
        musl=1.2.5-r8 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:a59bc827c80c4dbec2473c6b1a6f72b5b015d44e37981b112b9a60cb3c5180af

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
