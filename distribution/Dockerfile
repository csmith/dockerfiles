# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.45-r0","apk:libmagic":"5.45-r0","apk:make":"4.4.1-r1","apk:musl":"1.2.4-r1","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"59d70d137edf684ebef3dccd7bbcacc7ad173c4f2bae1fcd8b3a15e827e9a891"}

FROM reg.c5h.io/golang@sha256:59d70d137edf684ebef3dccd7bbcacc7ad173c4f2bae1fcd8b3a15e827e9a891 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.45-r0 \
        libmagic=5.45-r0 \
        make=4.4.1-r1 \
        musl=1.2.4-r1 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
