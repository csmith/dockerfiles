# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.42-r0","apk:g++":"13.2.1_git20240309-r0","apk:gcc":"13.2.1_git20240309-r0","apk:gmp":"6.3.0-r1","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20240309-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libgomp":"13.2.1_git20240309-r0","apk:libstdc++":"13.2.1_git20240309-r0","apk:libstdc++-dev":"13.2.1_git20240309-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r0","apk:musl-dev":"1.2.5-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:gohugoio/hugo":"v0.133.0","github:magefile/mage":"v1.15.0","image:base-glibc":"0fd47f2952d8787ae7164b111b85a739b3d9dac73858c7dcaea59d1cfa5f2f61","image:golang":"19d46582f163eefb8a4a6ac3106238f58921369736045b489e672c2e5d7d990a"}

FROM reg.c5h.io/golang@sha256:19d46582f163eefb8a4a6ac3106238f58921369736045b489e672c2e5d7d990a AS build

ARG TAG="v0.133.0"
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

FROM reg.c5h.io/base-glibc@sha256:0fd47f2952d8787ae7164b111b85a739b3d9dac73858c7dcaea59d1cfa5f2f61

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
