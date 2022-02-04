# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.2-r7","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968","image:golang":"ac8fa5f4078b0a697796b5d74158ab4e73b5f743433f621d89abb7c5faf94f1f"}

FROM reg.c5h.io/golang@sha256:ac8fa5f4078b0a697796b5d74158ab4e73b5f743433f621d89abb7c5faf94f1f AS build

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

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
