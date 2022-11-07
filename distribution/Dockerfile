# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.3-r1","image:base":"2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665","image:golang":"c943e10ea8acda82bf2d461bb2a083c59df867cb05f7e001d777c321bf9e08f0"}

FROM reg.c5h.io/golang@sha256:c943e10ea8acda82bf2d461bb2a083c59df867cb05f7e001d777c321bf9e08f0 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.41-r0 \
        libmagic=5.41-r0 \
        make=4.3-r0 \
        musl=1.2.3-r1 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
