# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.43.1-r2","apk:g++":"14.2.0-r4","apk:gcc":"14.2.0-r4","apk:gmp":"6.3.0-r2","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libstdc++":"14.2.0-r4","apk:libstdc++-dev":"14.2.0-r4","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r9","apk:musl-dev":"1.2.5-r9","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","github:gohugoio/hugo":"v0.145.0","github:magefile/mage":"v1.15.0","image:base-glibc":"700779c8a0724fe0d5945a552400724b2f68d9d02ed596c7d04140f1e7a355d1","image:golang":"993039fb677473d19819fc3f0448f020003b6917ac1a49a0605e5d50a4b15420"}

FROM reg.c5h.io/golang@sha256:993039fb677473d19819fc3f0448f020003b6917ac1a49a0605e5d50a4b15420 AS build

RUN set -eux; \
    apk add --no-cache \
        binutils=2.43.1-r2 \
        g++=14.2.0-r4 \
        gcc=14.2.0-r4 \
        gmp=6.3.0-r2 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=14.2.0-r4 \
        libgcc=14.2.0-r4 \
        libgomp=14.2.0-r4 \
        libstdc++=14.2.0-r4 \
        libstdc++-dev=14.2.0-r4 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r9 \
        musl-dev=1.2.5-r9 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.145.0 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.15.0"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:700779c8a0724fe0d5945a552400724b2f68d9d02ed596c7d04140f1e7a355d1

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
