# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.43-r0","apk:libmagic":"5.43-r0","apk:make":"4.3-r1","apk:musl":"1.2.3-r4","image:base":"c41c52eb39d01d876edd44db743514a484b09b61793b6ada6f064315eb18327d","image:golang":"da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e"}

FROM reg.c5h.io/golang@sha256:da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e AS build

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

FROM reg.c5h.io/base@sha256:c41c52eb39d01d876edd44db743514a484b09b61793b6ada6f064315eb18327d

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
