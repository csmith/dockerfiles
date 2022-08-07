# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.3-r0","image:base":"4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49","image:golang":"7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1"}

FROM reg.c5h.io/golang@sha256:7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.41-r0 \
        libmagic=5.41-r0 \
        make=4.3-r0 \
        musl=1.2.3-r0 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
