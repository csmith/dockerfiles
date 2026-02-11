# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.46-r2","apk:libmagic":"5.46-r2","apk:make":"4.4.1-r3","apk:musl":"1.2.5-r21","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"77752e4278043c67ff36b4335f475e90e765540450f428669249c4441cf287db"}

FROM reg.c5h.io/golang@sha256:77752e4278043c67ff36b4335f475e90e765540450f428669249c4441cf287db AS build

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

FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
