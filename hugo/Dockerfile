# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.43.1-r1","apk:g++":"14.2.0-r4","apk:gcc":"14.2.0-r4","apk:gmp":"6.3.0-r2","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libstdc++":"14.2.0-r4","apk:libstdc++-dev":"14.2.0-r4","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r8","apk:musl-dev":"1.2.5-r8","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r1","github:gohugoio/hugo":"v0.140.1","github:magefile/mage":"v1.15.0","image:base-glibc":"328a35e847326e90bb195327f52180b89f3b5599ba685b7da798a9333e4df5a5","image:golang":"5b55359c685b704a76e0ae30f538d9332e06a59d131ee57d2352ac214f3cdcd3"}

FROM reg.c5h.io/golang@sha256:5b55359c685b704a76e0ae30f538d9332e06a59d131ee57d2352ac214f3cdcd3 AS build

ARG TAG="v0.140.1"
ARG MAGE="v1.15.0"

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
        zstd-libs=1.5.6-r1 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:328a35e847326e90bb195327f52180b89f3b5599ba685b7da798a9333e4df5a5

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
