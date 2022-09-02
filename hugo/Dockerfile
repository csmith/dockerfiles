# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.38-r3","apk:g++":"11.2.1_git20220219-r2","apk:gcc":"11.2.1_git20220219-r2","apk:gmp":"6.2.1-r2","apk:isl22":"0.22-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libc-dev":"0.7.2-r3","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libstdc++":"11.2.1_git20220219-r2","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r0","apk:musl-dev":"1.2.3-r0","apk:zlib":"1.2.12-r3","github:gohugoio/hugo":"v0.102.3","github:magefile/mage":"v1.13.0","image:base-glibc":"3dafd2c2d6ecb6b31e527a103800fe846b1d98d5d2c4b8ecbeaa410153325fc8","image:golang":"290e92565f9e74b38a72d91bf56cca4c671b7e6de113b0ae98261b9f80c43d13"}

FROM reg.c5h.io/golang@sha256:290e92565f9e74b38a72d91bf56cca4c671b7e6de113b0ae98261b9f80c43d13 AS build

ARG TAG="v0.102.3"
ARG MAGE="v1.13.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.38-r3 \
        g++=11.2.1_git20220219-r2 \
        gcc=11.2.1_git20220219-r2 \
        gmp=6.2.1-r2 \
        isl22=0.22-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libc-dev=0.7.2-r3 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libstdc++=11.2.1_git20220219-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        zlib=1.2.12-r3 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:3dafd2c2d6ecb6b31e527a103800fe846b1d98d5d2c4b8ecbeaa410153325fc8

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
