# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r2","apk:libmagic":"5.46-r2","apk:make":"4.4.1-r3","apk:musl":"1.2.5-r23","image:base":"00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958","image:golang":"66b2fc993c20f1598d8f877bc6228c2f814d0d3c1055f744c1b1f3a3bc355023"}

FROM reg.c5h.io/golang@sha256:66b2fc993c20f1598d8f877bc6228c2f814d0d3c1055f744c1b1f3a3bc355023 AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r3 \
        musl=1.2.5-r23 \
        file=5.46-r2 \
        libmagic=5.46-r2 \
        musl=1.2.5-r23 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:00146a62a048620af617bc351091bce6105ca77ea883f65bf45aa22aecf62958

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
