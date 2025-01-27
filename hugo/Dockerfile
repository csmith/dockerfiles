# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.43.1-r1","apk:g++":"14.2.0-r4","apk:gcc":"14.2.0-r4","apk:gmp":"6.3.0-r2","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libstdc++":"14.2.0-r4","apk:libstdc++-dev":"14.2.0-r4","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r8","apk:musl-dev":"1.2.5-r8","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","github:gohugoio/hugo":"v0.142.0","github:magefile/mage":"v1.15.0","image:base-glibc":"e63442faf23c484fb79c173485b7a5a626b7296ed4c280d838f68d4a5f4cc21a","image:golang":"b6640cf9e9db2e2132ed4722ea3221bad060651fe0ed20803a17752d88c18c6a"}

FROM reg.c5h.io/golang@sha256:b6640cf9e9db2e2132ed4722ea3221bad060651fe0ed20803a17752d88c18c6a AS build

RUN set -eux; \
    apk add --no-cache \
        binutils=2.43.1-r1 \
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
        musl=1.2.5-r8 \
        musl-dev=1.2.5-r8 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.142.0 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.15.0"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:e63442faf23c484fb79c173485b7a5a626b7296ed4c280d838f68d4a5f4cc21a

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
