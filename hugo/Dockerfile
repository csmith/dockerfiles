# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.45.1-r0","apk:g++":"15.2.0-r2","apk:gcc":"15.2.0-r2","apk:gmp":"6.3.0-r4","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"15.2.0-r2","apk:libgcc":"15.2.0-r2","apk:libgomp":"15.2.0-r2","apk:libstdc++":"15.2.0-r2","apk:libstdc++-dev":"15.2.0-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.2-r0","apk:musl":"1.2.5-r21","apk:musl-dev":"1.2.5-r21","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.7-r2","github:gohugoio/hugo":"v0.154.5","github:magefile/mage":"v1.15.0","image:base-glibc":"83f9aed2018b2a72784d5a69df14786df3f6d93dc41f3795036d86a1aa55d5f2","image:golang":"039eff401799300a5fd431d8be0ae5955d2ec6f5393dde86e0b245a96f914189"}

FROM reg.c5h.io/golang@sha256:039eff401799300a5fd431d8be0ae5955d2ec6f5393dde86e0b245a96f914189 AS build

RUN set -eux; \
    apk add --no-cache \
        binutils=2.45.1-r0 \
        g++=15.2.0-r2 \
        gcc=15.2.0-r2 \
        gmp=6.3.0-r4 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=15.2.0-r2 \
        libgcc=15.2.0-r2 \
        libgomp=15.2.0-r2 \
        libstdc++=15.2.0-r2 \
        libstdc++-dev=15.2.0-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.2-r0 \
        musl=1.2.5-r21 \
        musl-dev=1.2.5-r21 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r2 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.154.5 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.15.0"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:83f9aed2018b2a72784d5a69df14786df3f6d93dc41f3795036d86a1aa55d5f2

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
