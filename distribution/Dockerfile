# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.41-r0","apk:libmagic":"5.41-r0","apk:make":"4.3-r0","apk:musl":"1.2.3-r0","image:base":"dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df","image:golang":"48b418f4f5f5a9c237b64ad04bd1f94b25ec0be5e06bfc29a37384ea450b4ac3"}

FROM reg.c5h.io/golang@sha256:48b418f4f5f5a9c237b64ad04bd1f94b25ec0be5e06bfc29a37384ea450b4ac3 AS build

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

FROM reg.c5h.io/base@sha256:dad97084335f003c2ec5251b8f9d22e59704f59c8e2ffcc6108470131341b2df

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
