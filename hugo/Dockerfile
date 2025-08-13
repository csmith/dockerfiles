# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.44-r2","apk:g++":"14.2.0-r6","apk:gcc":"14.2.0-r6","apk:gmp":"6.3.0-r3","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"14.2.0-r6","apk:libgcc":"14.2.0-r6","apk:libgomp":"14.2.0-r6","apk:libstdc++":"14.2.0-r6","apk:libstdc++-dev":"14.2.0-r6","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1_p1-r0","apk:musl":"1.2.5-r10","apk:musl-dev":"1.2.5-r10","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.7-r0","github:gohugoio/hugo":"v0.148.2","github:magefile/mage":"v1.15.0","image:base-glibc":"f4721399ab8b1bc5aa09b427008f6eaae9de0bbe92a4dd605b9be8ab1c29eb84","image:golang":"954b97f11d387f53f8aede05ec76d617607ceff094bfe9a84d19e208bec43f5e"}

FROM reg.c5h.io/golang@sha256:954b97f11d387f53f8aede05ec76d617607ceff094bfe9a84d19e208bec43f5e AS build

RUN set -eux; \
    apk add --no-cache \
        binutils=2.44-r2 \
        g++=14.2.0-r6 \
        gcc=14.2.0-r6 \
        gmp=6.3.0-r3 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=14.2.0-r6 \
        libgcc=14.2.0-r6 \
        libgomp=14.2.0-r6 \
        libstdc++=14.2.0-r6 \
        libstdc++-dev=14.2.0-r6 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1_p1-r0 \
        musl=1.2.5-r10 \
        musl-dev=1.2.5-r10 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.148.2 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.15.0"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:f4721399ab8b1bc5aa09b427008f6eaae9de0bbe92a4dd605b9be8ab1c29eb84

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
