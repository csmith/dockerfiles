# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.39-r2","apk:g++":"12.2.1_git20220924-r4","apk:gcc":"12.2.1_git20220924-r4","apk:gmp":"6.2.1-r2","apk:isl24":"0.24-r1","apk:libatomic":"12.2.1_git20220924-r4","apk:libc-dev":"0.7.2-r3","apk:libgcc":"12.2.1_git20220924-r4","apk:libgomp":"12.2.1_git20220924-r4","apk:libstdc++":"12.2.1_git20220924-r4","apk:libstdc++-dev":"12.2.1_git20220924-r4","apk:mpc1":"1.2.1-r1","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r4","apk:musl-dev":"1.2.3-r4","apk:zlib":"1.2.13-r0","github:gohugoio/hugo":"v0.107.0","github:magefile/mage":"v1.14.0","image:base-glibc":"f48a150b0d278d1c00173ce0129f3845fa86a3071f3b2db045c1a47b2c3fdacd","image:golang":"55f91d869bdbd7a6c800e9095e5a0596e359db56beb7aad88697ff488b25c994"}

FROM reg.c5h.io/golang@sha256:55f91d869bdbd7a6c800e9095e5a0596e359db56beb7aad88697ff488b25c994 AS build

ARG TAG="v0.107.0"
ARG MAGE="v1.14.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.39-r2 \
        g++=12.2.1_git20220924-r4 \
        gcc=12.2.1_git20220924-r4 \
        gmp=6.2.1-r2 \
        isl24=0.24-r1 \
        libatomic=12.2.1_git20220924-r4 \
        libc-dev=0.7.2-r3 \
        libgcc=12.2.1_git20220924-r4 \
        libgomp=12.2.1_git20220924-r4 \
        libstdc++=12.2.1_git20220924-r4 \
        libstdc++-dev=12.2.1_git20220924-r4 \
        mpc1=1.2.1-r1 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r4 \
        musl-dev=1.2.3-r4 \
        zlib=1.2.13-r0 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:f48a150b0d278d1c00173ce0129f3845fa86a3071f3b2db045c1a47b2c3fdacd

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
