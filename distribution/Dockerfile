# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r0","apk:libmagic":"5.46-r0","apk:make":"4.4.1-r2","apk:musl":"1.2.5-r8","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"f11bab806d7b9d164d3d52afa3ac5c0f976cb5b78a09c1986ed6d25c971c2a45"}

FROM reg.c5h.io/golang@sha256:f11bab806d7b9d164d3d52afa3ac5c0f976cb5b78a09c1986ed6d25c971c2a45 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.46-r0 \
        libmagic=5.46-r0 \
        make=4.4.1-r2 \
        musl=1.2.5-r8 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
