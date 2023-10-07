# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.45-r0","apk:libmagic":"5.45-r0","apk:make":"4.4.1-r1","apk:musl":"1.2.4-r2","image:base":"7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f","image:golang":"8ec62e029f2954b79051926ca2270e2bec6e1ac887188593ea0b366c826cbbc0"}

FROM reg.c5h.io/golang@sha256:8ec62e029f2954b79051926ca2270e2bec6e1ac887188593ea0b366c826cbbc0 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.45-r0 \
        libmagic=5.45-r0 \
        make=4.4.1-r1 \
        musl=1.2.4-r2 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
