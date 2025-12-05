# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r2","apk:libmagic":"5.46-r2","apk:make":"4.4.1-r3","apk:musl":"1.2.5-r21","image:base":"46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7","image:golang":"18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef"}

FROM reg.c5h.io/golang@sha256:18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r3 \
        musl=1.2.5-r21 \
        file=5.46-r2 \
        libmagic=5.46-r2 \
        musl=1.2.5-r21 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:46ea216adf9eb9615e03305be16d497ef5e1f3d936f3e986a8e4d25e599f27f7

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
