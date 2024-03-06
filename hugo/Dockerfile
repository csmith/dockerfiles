# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.41-r0","apk:g++":"13.2.1_git20231014-r0","apk:gcc":"13.2.1_git20231014-r0","apk:gmp":"6.3.0-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20231014-r0","apk:libc-dev":"0.7.2-r5","apk:libgcc":"13.2.1_git20231014-r0","apk:libgomp":"13.2.1_git20231014-r0","apk:libstdc++":"13.2.1_git20231014-r0","apk:libstdc++-dev":"13.2.1_git20231014-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.4_git20230717-r4","apk:musl-dev":"1.2.4_git20230717-r4","apk:zlib":"1.3.1-r0","apk:zstd-libs":"1.5.5-r8","github:gohugoio/hugo":"v0.123.7","github:magefile/mage":"v1.15.0","image:base-glibc":"74542299930c372a26a4751088b4b22564070c9c28bc3eb2960df2905a7708b3","image:golang":"76b2d33b8addb686992b9154a584b8b7e0c3dd8b7c4ed80c9813c04d54e9ef74"}

FROM reg.c5h.io/golang@sha256:76b2d33b8addb686992b9154a584b8b7e0c3dd8b7c4ed80c9813c04d54e9ef74 AS build

ARG TAG="v0.123.7"
ARG MAGE="v1.15.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.41-r0 \
        g++=13.2.1_git20231014-r0 \
        gcc=13.2.1_git20231014-r0 \
        gmp=6.3.0-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20231014-r0 \
        libc-dev=0.7.2-r5 \
        libgcc=13.2.1_git20231014-r0 \
        libgomp=13.2.1_git20231014-r0 \
        libstdc++=13.2.1_git20231014-r0 \
        libstdc++-dev=13.2.1_git20231014-r0 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.4_git20230717-r4 \
        musl-dev=1.2.4_git20230717-r4 \
        zlib=1.3.1-r0 \
        zstd-libs=1.5.5-r8 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:74542299930c372a26a4751088b4b22564070c9c28bc3eb2960df2905a7708b3

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
