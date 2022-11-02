# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.3-r0","image:base":"893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1","image:golang":"16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f"}

FROM reg.c5h.io/golang@sha256:16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f AS build

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

FROM reg.c5h.io/base@sha256:893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
