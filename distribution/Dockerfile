# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.45-r1","apk:libmagic":"5.45-r1","apk:make":"4.4.1-r2","apk:musl":"1.2.5-r0","image:base":"b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff","image:golang":"34ab217721b830f8756683c8b590b06355aa4c33255f04a10e620fa41b07c72b"}

FROM reg.c5h.io/golang@sha256:34ab217721b830f8756683c8b590b06355aa4c33255f04a10e620fa41b07c72b AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.45-r1 \
        libmagic=5.45-r1 \
        make=4.4.1-r2 \
        musl=1.2.5-r0 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:b92ffcd70a3cea70972a600a917a409451afe2732844d09e3afe212ad8e041ff

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
