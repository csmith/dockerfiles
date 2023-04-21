# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.43-r0","apk:libmagic":"5.43-r0","apk:make":"4.3-r1","apk:musl":"1.2.3-r4","image:base":"d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092","image:golang":"00fba4baba7f5753b82ad1e03546cd02622bad16b6dc6f85bd9a79e2189f8cc3"}

FROM reg.c5h.io/golang@sha256:00fba4baba7f5753b82ad1e03546cd02622bad16b6dc6f85bd9a79e2189f8cc3 AS build

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

FROM reg.c5h.io/base@sha256:d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
