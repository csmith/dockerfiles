# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.45-r1","apk:libmagic":"5.45-r1","apk:make":"4.4.1-r2","apk:musl":"1.2.4_git20230717-r4","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"0304828501e1f9414244d2b381c5a8276ade33eb0a6c89bbce9802fcfab2ce4c"}

FROM reg.c5h.io/golang@sha256:0304828501e1f9414244d2b381c5a8276ade33eb0a6c89bbce9802fcfab2ce4c AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.45-r1 \
        libmagic=5.45-r1 \
        make=4.4.1-r2 \
        musl=1.2.4_git20230717-r4 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
