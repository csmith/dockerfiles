# Generated from https://github.com/csmith/dockerfiles/blob/master/hugo/Dockerfile.gotpl
# BOM: {"apk:binutils":"2.41-r0","apk:g++":"13.2.1_git20231014-r0","apk:gcc":"13.2.1_git20231014-r0","apk:gmp":"6.3.0-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20231014-r0","apk:libc-dev":"0.7.2-r5","apk:libgcc":"13.2.1_git20231014-r0","apk:libgomp":"13.2.1_git20231014-r0","apk:libstdc++":"13.2.1_git20231014-r0","apk:libstdc++-dev":"13.2.1_git20231014-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.4_git20230717-r4","apk:musl-dev":"1.2.4_git20230717-r4","apk:zlib":"1.3.1-r0","apk:zstd-libs":"1.5.5-r8","github:gohugoio/hugo":"v0.124.0","github:magefile/mage":"v1.15.0","image:base-glibc":"00f789a41d781840d9be09031040b438f203a52240abe1f6baa7d3f50df77ff6","image:golang":"02f9a78797fac2b87fbd32a53f473fdc1c0ef10b68ceb3deb21436cf9f069e7e"}

FROM reg.c5h.io/golang@sha256:02f9a78797fac2b87fbd32a53f473fdc1c0ef10b68ceb3deb21436cf9f069e7e AS build

ARG TAG="v0.124.0"
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

FROM reg.c5h.io/base-glibc@sha256:00f789a41d781840d9be09031040b438f203a52240abe1f6baa7d3f50df77ff6

COPY --from=build /go/bin/hugo /usr/bin/hugo
COPY --from=build --chown=65532:65532 /tmp/site /
VOLUME /site
WORKDIR /site
ENTRYPOINT ["/usr/bin/hugo"]
