# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.2-r7","image:base":"e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903","image:golang":"72adae738352aa1e5dc312952446651fc767661e2b9bdb006be122bfde7fd03a"}

FROM reg.c5h.io/golang@sha256:72adae738352aa1e5dc312952446651fc767661e2b9bdb006be122bfde7fd03a AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.41-r0 \
        libmagic=5.41-r0 \
        make=4.3-r0 \
        musl=1.2.2-r7 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /go/src/github.com/docker/distribution/bin/registry;

FROM reg.c5h.io/base@sha256:e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
