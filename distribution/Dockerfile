# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r0","apk:libmagic":"5.46-r0","apk:make":"4.4.1-r2","apk:musl":"1.2.5-r8","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"5b55359c685b704a76e0ae30f538d9332e06a59d131ee57d2352ac214f3cdcd3"}

FROM reg.c5h.io/golang@sha256:5b55359c685b704a76e0ae30f538d9332e06a59d131ee57d2352ac214f3cdcd3 AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r2 \
        musl=1.2.5-r8 \
        file=5.46-r0 \
        libmagic=5.46-r0 \
        musl=1.2.5-r8 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
