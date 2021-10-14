# Generated from https://github.com/csmith/dockerfiles/blob/master/distribution/Dockerfile.gotpl
# BOM: {"apk:file":"5.40-r1","apk:libmagic":"5.40-r1","apk:make":"4.3-r0","apk:musl":"1.2.2-r3","image:base":"sha256:6c8430c1a55197f17620bd0933b5cecc648b0bcb56ef60d6f97c766e0ecbb437","image:golang":"sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670"}

FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v2.7.1"

RUN set -eux; \
    apk add --no-cache \
        file=5.40-r1 \
        libmagic=5.40-r1 \
        make=4.3-r0 \
        musl=1.2.2-r3 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/distribution/distribution /go/src/github.com/docker/distribution; \
    cd /go/src/github.com/docker/distribution; \
    BUILDTAGS="include_oss include_gcs" GO111MODULE=off CGO_ENABLED=0 make bin/registry; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /go/src/github.com/docker/distribution/bin/registry;

FROM reg.c5h.io/base@sha256:6c8430c1a55197f17620bd0933b5cecc648b0bcb56ef60d6f97c766e0ecbb437

COPY --from=build /go/src/github.com/docker/distribution/cmd/registry/config-dev.yml /etc/docker/registry/config.yml
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
EXPOSE 5000
VOLUME ["/var/lib/registry"]
ENTRYPOINT ["/bin/registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
