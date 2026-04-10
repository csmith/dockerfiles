# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.45.1-r0","apk:g++":"15.2.0-r2","apk:gcc":"15.2.0-r2","apk:gmp":"6.3.0-r4","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"15.2.0-r2","apk:libgcc":"15.2.0-r2","apk:libgomp":"15.2.0-r2","apk:libstdc++":"15.2.0-r2","apk:libstdc++-dev":"15.2.0-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.2-r0","apk:musl":"1.2.5-r22","apk:musl-dev":"1.2.5-r22","apk:zlib":"1.3.2-r0","apk:zstd-libs":"1.5.7-r2","github:gohugoio/hugo":"v0.160.1","github:magefile/mage":"v1.17.1","image:base-glibc":"ffcdea4b139048a31085755cb12a28448a039cdf92fe6b0d7319e9378c5c42cb","image:golang":"d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a"}

FROM reg.c5h.io/golang@sha256:d2c7d4b907467670c9dc335754a890a061167ede406e99020e7d7e1a43eada6a AS build

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
        musl=1.2.5-r22 \
        musl-dev=1.2.5-r22 \
        zlib=1.3.2-r0 \
        zstd-libs=1.5.7-r2 \
    ;


RUN set -eux; \
    git clone --depth=1 -b v0.160.1 --single-branch https://github.com/gohugoio/hugo /tmp/work

WORKDIR /tmp/work


ARG MAGE="v1.17.1"

RUN set -eux; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:ffcdea4b139048a31085755cb12a28448a039cdf92fe6b0d7319e9378c5c42cb

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
