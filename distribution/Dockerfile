# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.40-r1","apk:libmagic":"5.40-r1","apk:make":"4.3-r0","apk:musl":"1.2.2-r3","image:base":"sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53","image:golang":"sha256:c6f396c63f87b6463b8090b96a2c9eea9c6065d52a2621cbc4e61f7a6b1c8128"}

FROM reg.c5h.io/golang@sha256:c6f396c63f87b6463b8090b96a2c9eea9c6065d52a2621cbc4e61f7a6b1c8128 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.40-r1 \
        libmagic=5.40-r1 \
        make=4.3-r0 \
        musl=1.2.2-r3 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /go/src/github.com/docker/distribution/bin/registry;

FROM reg.c5h.io/base@sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
