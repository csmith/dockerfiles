# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.43-r0","apk:libmagic":"5.43-r0","apk:make":"4.3-r1","apk:musl":"1.2.3-r4","image:base":"d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743","image:golang":"060342932f91f2d3334c4b309d6ffbb7482e63332e6385dd1c97919074141cb8"}

FROM reg.c5h.io/golang@sha256:060342932f91f2d3334c4b309d6ffbb7482e63332e6385dd1c97919074141cb8 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.43-r0 \
        libmagic=5.43-r0 \
        make=4.3-r1 \
        musl=1.2.3-r4 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
