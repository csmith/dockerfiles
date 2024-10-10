# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.42-r0","apk:g++":"13.2.1_git20240309-r0","apk:gcc":"13.2.1_git20240309-r0","apk:gmp":"6.3.0-r1","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20240309-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libgomp":"13.2.1_git20240309-r0","apk:libstdc++":"13.2.1_git20240309-r0","apk:libstdc++-dev":"13.2.1_git20240309-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r0","apk:musl-dev":"1.2.5-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:gohugoio/hugo":"v0.135.0","github:magefile/mage":"v1.15.0","image:base-glibc":"edc8b4c11cc6d577cf738171228de8753242d1ef51b9c61a5df964b66c432112","image:golang":"287c2dc187f952ecfdbecb0b5cb612a9faa4d7ebfbb0bee3f4021172fbac1971"}

FROM reg.c5h.io/golang@sha256:287c2dc187f952ecfdbecb0b5cb612a9faa4d7ebfbb0bee3f4021172fbac1971 AS build

ARG TAG="v0.135.0"
ARG MAGE="v1.15.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.42-r0 \
        g++=13.2.1_git20240309-r0 \
        gcc=13.2.1_git20240309-r0 \
        gmp=6.3.0-r1 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20240309-r0 \
        libgcc=13.2.1_git20240309-r0 \
        libgomp=13.2.1_git20240309-r0 \
        libstdc++=13.2.1_git20240309-r0 \
        libstdc++-dev=13.2.1_git20240309-r0 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r0 \
        musl-dev=1.2.5-r0 \
        zlib=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:edc8b4c11cc6d577cf738171228de8753242d1ef51b9c61a5df964b66c432112

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
