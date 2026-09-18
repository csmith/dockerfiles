# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.47-r2","apk:libmagic":"5.47-r2","apk:make":"4.4.1-r4","apk:musl":"1.2.6-r2","image:base":"68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc","image:golang":"424727a3913a09b7b5f3aa8b43cc60a6aded1a6b814d2ca889f7305c2a06c6a3"}

FROM reg.c5h.io/golang@sha256:424727a3913a09b7b5f3aa8b43cc60a6aded1a6b814d2ca889f7305c2a06c6a3 AS build

RUN set -eux; \
    apk add --no-cache \
        make=4.4.1-r4 \
        musl=1.2.6-r2 \
        file=5.47-r2 \
        libmagic=5.47-r2 \
        musl=1.2.6-r2 \
    ;

RUN set -eux; \
    git clone --depth=1 -b v2.7.1 --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution

WORKDIR /go/src/github.com/docker/distribution


RUN set -eux; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:68804ff33ba310905be47653301a7636727a352c4b820d9e5e672df96577a3dc

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
