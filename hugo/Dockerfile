# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.40-r7","apk:g++":"12.2.1_git20220924-r10","apk:gcc":"12.2.1_git20220924-r10","apk:gmp":"6.2.1-r3","apk:isl25":"0.25-r2","apk:libatomic":"12.2.1_git20220924-r10","apk:libc-dev":"0.7.2-r5","apk:libgcc":"12.2.1_git20220924-r10","apk:libgomp":"12.2.1_git20220924-r10","apk:libstdc++":"12.2.1_git20220924-r10","apk:libstdc++-dev":"12.2.1_git20220924-r10","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.0_p12-r0","apk:musl":"1.2.4-r1","apk:musl-dev":"1.2.4-r1","apk:zlib":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","github:gohugoio/hugo":"v0.117.0","github:magefile/mage":"v1.15.0","image:base-glibc":"452263727e2fae99208974931dda91ad027be7016f5ee99fe74134ab158b4320","image:golang":"3e8330fcd48f7f61c9f6f00b1b9bb450c848225f37e39cb00efab8c0a1756470"}

FROM reg.c5h.io/golang@sha256:3e8330fcd48f7f61c9f6f00b1b9bb450c848225f37e39cb00efab8c0a1756470 AS build

ARG TAG="v0.117.0"
ARG MAGE="v1.15.0"

RUN set -eux; \
    apk add --no-cache \
        binutils=2.40-r7 \
        g++=12.2.1_git20220924-r10 \
        gcc=12.2.1_git20220924-r10 \
        gmp=6.2.1-r3 \
        isl25=0.25-r2 \
        libatomic=12.2.1_git20220924-r10 \
        libc-dev=0.7.2-r5 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libstdc++=12.2.1_git20220924-r10 \
        libstdc++-dev=12.2.1_git20220924-r10 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0_p12-r0 \
        musl=1.2.4-r1 \
        musl-dev=1.2.4-r1 \
        zlib=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
        ; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/gohugoio/hugo /tmp/work; \
    cd /tmp/work; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE hugo; \
    HUGO_BUILD_TAGS=extended go run github.com/magefile/mage@$MAGE install; \
    mkdir -p /tmp/site;

FROM reg.c5h.io/base-glibc@sha256:452263727e2fae99208974931dda91ad027be7016f5ee99fe74134ab158b4320

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
