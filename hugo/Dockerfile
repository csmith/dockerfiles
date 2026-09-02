# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.45.1-r1","apk:g++":"15.2.0-r5","apk:gcc":"15.2.0-r5","apk:gmp":"6.3.0-r4","apk:isl25":"0.25-r2","apk:jansson":"2.15.0-r0","apk:libatomic":"15.2.0-r5","apk:libgcc":"15.2.0-r5","apk:libgcc-static":"15.2.0-r5","apk:libgomp":"15.2.0-r5","apk:libstdc++":"15.2.0-r5","apk:libstdc++-dev":"15.2.0-r5","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.2-r0","apk:musl":"1.2.6-r2","apk:musl-dev":"1.2.6-r2","apk:zlib":"1.3.2-r0","apk:zstd-libs":"1.5.7-r2","github:https://github.com/gohugoio/hugo":"v0.165.0","github:https://github.com/magefile/mage":"v1.17.2","image:base-glibc":"4456658d241ef53a8b27b3052bf69fca1d0c4766d97b0c360059419ecb3d539d","image:golang":"3204d08f9b5525b3641537d5e1159e8e751622811967f646de0e108b8148fa59"}

FROM reg.c5h.io/golang@sha256:3204d08f9b5525b3641537d5e1159e8e751622811967f646de0e108b8148fa59 AS build

RUN set -eux; \
    apk add --no-cache \
        binutils=2.45.1-r1 \
        g++=15.2.0-r5 \
        gcc=15.2.0-r5 \
        gmp=6.3.0-r4 \
        isl25=0.25-r2 \
        jansson=2.15.0-r0 \
        libatomic=15.2.0-r5 \
        libgcc=15.2.0-r5 \
        libgcc-static=15.2.0-r5 \
        libgomp=15.2.0-r5 \
        libstdc++=15.2.0-r5 \
        libstdc++-dev=15.2.0-r5 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.2-r0 \
        musl=1.2.6-r2 \
        musl-dev=1.2.6-r2 \
        zlib=1.3.2-r0 \
        zstd-libs=1.5.7-r2 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.165.0 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.17.2"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:4456658d241ef53a8b27b3052bf69fca1d0c4766d97b0c360059419ecb3d539d

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
